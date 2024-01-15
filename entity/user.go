package entity

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"-"`
	Age      int    `json:"age"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
