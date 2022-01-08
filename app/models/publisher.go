package models

type Publisher struct {
	ID     int32  `json:"id" gorm:"primary_key"`
	BookID int32  `json:"bookID"`
	Name   string `json:"name"`
}
