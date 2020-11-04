package persistence

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"{{.Module}}/option"
	xzl "github.com/jacexh/xorm-zap-logger"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

var (
	db *xorm.Engine
)

// BuildDBConnection 数据库连接函数
func BuildDBConnection(option option.DatabaseOption, logger *zap.Logger) (*xorm.Engine, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		option.Username, option.Password, option.Host, option.Port, option.Database)
	var err error
	db, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(option.MaxIdleConnections)
	db.SetMaxOpenConns(option.MaxOpenConnections)

	db.SetLogger(xzl.NewXormZapLogger(logger))
	db.ShowSQL(true)

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
