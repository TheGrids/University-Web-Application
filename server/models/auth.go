package models

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	EmailCheck bool   `json:"email_check"`
}

type UserRegisterData struct {
	Email     string `json:"email" binding:"required,email" gorm:"unique"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type EmailCheck struct {
	ID   uint   `json:"id"`
	UUID string `json:"uuid"`
}
