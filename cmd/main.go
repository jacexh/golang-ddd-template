package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jacexh/golang-ddd-template/internal/application"
	"github.com/jacexh/golang-ddd-template/internal/infrastructure/persistence"
	"github.com/jacexh/golang-ddd-template/internal/logger"
	"github.com/jacexh/golang-ddd-template/internal/option"
	"github.com/jacexh/golang-ddd-template/internal/transport/rest"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"go.uber.org/zap"
)

var (
	// version 项目版本号，可以在构建时传入
	version = "(git commit revision)"
)

func main() {
	// 加载项目配置文件
	opt := option.MustLoadConfig()

	// 加载全局日志配置，完成日志的初始化操作
	log := logger.BuildLogger(opt.Logger)
	logger.Logger.Info("loaded options", zap.Any("option", opt), zap.String("version", version))
	logger.SetTracer(&rest.ChiRequestIDTracer{})

	// 创建数据库连接
	db, err := persistence.BuildDBConnection(opt.Database, log)
	if err != nil {
		logger.Logger.Fatal("failed to connect with database", zap.Error(err))
	}
	ur := persistence.BuildUserRepository(db)

	// 初始化application层
	application.BuildUserApplication(ur)

	// 启动运行web server
	eng := rest.BuildRouter(opt.Router, log)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(opt.Router.Port),
		Handler: eng,
	}
	logger.Logger.Info(fmt.Sprintf("service is running on port %d", opt.Router.Port))

	errChan := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- srv.ListenAndServe()
		}
	}()

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-sigs
		err := fmt.Errorf("caught signal: %s", sig.String())
		logger.Logger.Warn("caught quit signal, try to shutdown service in 5 seconds", zap.String("signal", sig.String()))
		if e := srv.Shutdown(infection.GenContextWithTimeout(5 * time.Second)); e != nil {
			err = fmt.Errorf("%w | failed to shutdown: %s", err, e.Error())
		}
		errChan <- err
	}()

	err = <-errChan
	logger.Logger.Warn("kill all available contexts after 30 seconds")
	infection.KillContextsAfter(30 * time.Second)
	logger.Logger.Warn("service was shutdown", zap.Error(err))
	os.Exit(0)
}
