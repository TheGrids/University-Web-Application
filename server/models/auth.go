package models

type User struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Role            string `json:"role"`
	Faculty         string `json:"faculty" gorm:"default:'Нет информации'"`
	Group           string `json:"group" gorm:"default:'Нет информации'"`
	FormOfEducation string `json:"form_of_education" gorm:"default:'Нет информации'"`
	Level           string `json:"level" gorm:"default:'Нет информации'"`
	EmailCheck      bool   `json:"email_check"`
}

type UserRegisterData struct {
	Email     string `json:"email" binding:"required,email" gorm:"unique"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type UserLoginData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EmailCheck struct {
	ID   uint   `json:"id"`
	UUID string `json:"uuid"`
}

type Token struct {
	ID      uint   `json:"id"`
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
