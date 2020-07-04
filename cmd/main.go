package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"{{.Module}}/application"
	"{{.Module}}/infrastructure/repository"
	"{{.Module}}/logger"
	"{{.Module}}/option"
	"{{.Module}}/router"
	"github.com/jacexh/multiconfig"
	"go.uber.org/zap"
)

var (
	// version 项目版本号，可以在构建时传入
	version = "(git commit revision)"
	// environmentVariablesPrefix 项目环境变量前缀
	environmentVariablesPrefix = "{{.EnvironmentVariablesPrefix}}"
	// environmentVariableProfile 项目profile的环境变量名称
	environmentVariableProfile = environmentVariablesPrefix + "_PROJECT_PROFILE"
	// profileDirectoryPath 项目配置目录路径
	profileDirectoryPath = "."
	// profileFilePrefix profile文件前缀
	profileFilePrefix = "config"
	// profileFileFormat profile文件格式
	profileFileFormat = "yaml"
)

func loadOptionByProfile() *option.Option {
	profile := os.Getenv(environmentVariableProfile)
	fn := profileFilePrefix
	if profile != "" {
		fn = fn + "_" + strings.ToLower(profile)
	}
	var fs []string

	switch profileFileFormat {
	case "json":
		fs = append(fs, fn+".json")
	case "toml":
		fs = append(fs, fn+".toml")
	case "yaml":
		fs = append(fs, fn+".yml", fn+".yaml")
	default:
		panic(errors.New("unsupported file format"))
	}

	var err error
	_, runningFile, _, _ := runtime.Caller(1)
	for _, f := range fs {
		fp := filepath.Join(path.Dir(runningFile), profileDirectoryPath, f)
		_, err = os.Stat(fp)
		if os.IsNotExist(err) {
			continue
		}
		return loadOption(fp)
	}
	panic(err)
}

func loadOption(path string) *option.Option {
	loader := multiconfig.NewWithPathAndEnvPrefix(path, environmentVariablesPrefix)

	opt := new(option.Option)
	loader.MustLoad(opt)
	return opt
}

func main() {
	// 加载项目配置文件
	opt := loadOptionByProfile()

	// 加载全局日志配置，完成日志的初始化操作
	logger.BuildLogger(opt.Logger)
	logger.Logger.Info("loaded options", zap.Any("option", opt), zap.String("version", version))

	// 创建数据库连接
	db, err := repository.BuildDBConnection(opt.Database)
	if err != nil {
		logger.Logger.Panic("failed to connect with database", zap.Error(err))
	}
	ur := repository.NewUserRepository(db)

	// 初始化application层
	application.BuildUserApplication(ur)

	// 启动运行web server
	eng := router.BuildRouter(opt.Router)
	eng.GET("/ping", router.Ping)
	eng.GET("/users/:user", router.GetUser)

	// 服务启动
	errChan := make(chan error, 1)
	go func() {
		logger.Logger.Info(fmt.Sprintf("service is running on port %d", opt.Router.Port))
		errChan <- eng.Run(":" + strconv.Itoa(opt.Router.Port))
	}()

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-sigs
		errChan <- fmt.Errorf("caught signal: %s", sig.String())
	}()

	logger.Logger.Panic("service was shutdown", zap.Error(<-errChan))
}
