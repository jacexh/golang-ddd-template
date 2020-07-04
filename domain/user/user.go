package user

import "time"

type UserEntity struct {
	ID    string    `ddb:"id"`
	Name  string    `ddb:"name"`
	Email string    `ddb:"email"`
	CTime time.Time `ddb:"ctime"`
	MTime time.Time `ddb:"mtime"`
}
