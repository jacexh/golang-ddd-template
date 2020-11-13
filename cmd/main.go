package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"{{.Module}}/application"
	"{{.Module}}/infrastructure/persistence"
	"{{.Module}}/logger"
	"{{.Module}}/option"
	"{{.Module}}/router"
	"go.uber.org/zap"
)

var (
	// version 项目版本号，可以在构建时传入
	version = "(git commit revision)"
)

func main() {
	// 加载项目配置文件
	opt := option.LoadConfig()

	// 加载全局日志配置，完成日志的初始化操作
	log := logger.BuildLogger(opt.Logger)
	logger.Logger.Info("loaded options", zap.Any("option", opt), zap.String("version", version))

	// 创建数据库连接
	db, err := persistence.BuildDBConnection(opt.Database, log)
	if err != nil {
		logger.Logger.Fatal("failed to connect with database", zap.Error(err))
	}
	ur := persistence.BuildUserRepository(db)

	// 初始化application层
	application.BuildUserApplication(ur)

	// 启动运行web server
	eng := router.BuildRouter(opt.Router)

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

	logger.Logger.Fatal("service was shutdown", zap.Error(<-errChan))
}
