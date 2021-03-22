package dto

type (
	// UserDTO use的dto定义，用于api层
	UserDTO struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Version string `json:"version"`
		CTime   string `json:"ctime"`
		MTime   string `json:"mtime"`
	}
)
