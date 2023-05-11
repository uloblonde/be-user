package models

type User struct {
	Id       int
	Name     string `json:"name" gorm:"name"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}
