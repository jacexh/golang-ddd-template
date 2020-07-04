package main

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/repository"
	"github.com/jacexh/golang-ddd-template/router"
	"github.com/jacexh/golang-ddd-template/types"
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
	profileDirectoryPath = "./conf"
	// profileFilePrefix profile文件前缀
	profileFilePrefix = "config"
	// profileFileFormat profile文件格式
	profileFileFormat = "yaml"
)

func loadOptionByProfile() *types.Option {
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
	for _, f := range fs {
		fp := filepath.Join(profileDirectoryPath, f)
		_, err = os.Stat(fp)
		if os.IsNotExist(err) {
			continue
		}
		return loadOption(fp)
	}
	panic(err)
}

func loadOption(path string) *types.Option {
	loader := multiconfig.NewWithPathAndEnvPrefix(path, environmentVariablesPrefix)

	opt := new(types.Option)
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
	_, err := repository.BuildDBConnection(opt.Database)
	if err != nil {
		logger.Logger.Panic("failed to connect with database", zap.Error(err))
	}

	// 启动运行web server
	eng := router.BuildRouter(opt.Router)
	logger.Logger.Info("server is booting up")
	logger.Logger.Fatal(
		"server was down",
		zap.Error(eng.Run(":"+strconv.Itoa(opt.Router.Port))))
}
