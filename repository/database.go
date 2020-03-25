package repository

import (
	"database/sql"
	"sync"

	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jacexh/golang-ddd-template/types"
)

var (
	db   *sql.DB
	once sync.Once
)

// BuildDBConnection 数据库连接函数
func BuildDBConnection(option types.DatabaseOption) (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = manager.New(option.Database, option.Username, option.Password, option.Host).Set(
			manager.SetCharset("utf8mb4"),
			manager.SetParseTime(true),
			manager.SetAllowCleartextPasswords(true),
		).Port(option.Port).Open(true)
		if err != nil {
			return
		}
		db.SetMaxOpenConns(option.MaxOpenConnections)
		db.SetMaxIdleConns(option.MaxIdleConnections)
	})
	return db, err
}
