package do

import (
	"database/sql"
	"time"
)

type UserDo struct {
	ID       string       `ddb:"id"`
	Name     string       `ddb:"name"`
	Password sql.RawBytes `ddb:"password"`
	Email    string       `ddb:"email"`
	Version  int          `ddb:"version"`
	CTime    time.Time    `ddb:"ctime"`
	MTime    time.Time    `ddb:"mtime"`
}
