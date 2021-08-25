package dto

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}
