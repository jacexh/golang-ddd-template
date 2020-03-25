package entity

// User user数据库实体
type User struct {
	ID    string `json:"id" ddb:"id"`
	Name  string `json:"name" ddb:"name"`
	Email string `json:"email" ddb:"email"`
}
