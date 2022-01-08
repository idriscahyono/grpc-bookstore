package models

type Book struct {
	ID         int32  `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publishers []Publisher
}
