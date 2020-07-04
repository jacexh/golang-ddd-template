package repository

import (
	"database/sql"

	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"{{.Module}}/option"
)

var (
	db *sql.DB
)

// BuildDBConnection 数据库连接函数
func BuildDBConnection(option option.DatabaseOption) (*sql.DB, error) {
	var err error
	db, err = manager.New(option.Database, option.Username, option.Password, option.Host).Set(
		manager.SetCharset("utf8mb4"),
		manager.SetParseTime(true),
		manager.SetAllowCleartextPasswords(true),
	).Port(option.Port).Open(true)
	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(option.MaxOpenConnections)
	db.SetMaxIdleConns(option.MaxIdleConnections)
	return db, err
}
