package models

type Message struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	AuthorId uint   `json:"authorId"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Time     int64  `json:"time"`
}

type AddMessage struct {
	AuthorId uint   `json:"authorId" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Body     string `json:"body" binding:"required"`
}
