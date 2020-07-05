package do

import "time"

type UserDo struct {
	ID       string    `ddb:"id"`
	Name     string    `ddb:"name"`
	Password string    `ddb:"password"`
	Email    string    `ddb:"email"`
	Version  int       `ddb:"version"`
	CTime    time.Time `ddb:"ctime"`
	MTime    time.Time `ddb:"mtime"`
}
