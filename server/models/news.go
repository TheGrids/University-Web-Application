package models

type News struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	AuthorId        uint   `json:"authorId"`
	AuthorFirstName string `json:"authorFirstName"`
	AuthorLastName  string `json:"authorLastName"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	Tag             string `json:"tag"`
	Time            int64  `json:"time"`
}

type AddNews struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Tag   string `json:"tag" binding:"required"`
}
