package models

type GetProfile struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Role            string `json:"role"`
	Faculty         string `json:"faculty" gorm:"default:'Нет информации'"`
	Group           string `json:"group" gorm:"default:'Нет информации'"`
	FormOfEducation string `json:"form_of_education" gorm:"default:'Нет информации'"`
	Level           string `json:"level" gorm:"default:'Нет информации'"`
}
