package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/idriscahyono/grpc-bookstore/app/models"
	"gorm.io/gorm"
	"log"
)

func PublisherFaker(db *gorm.DB) *models.Publisher {
	book := BookFaker(db)
	err := db.Create(&book).Error
	if err != nil {
		log.Fatal(err)
	}
	return &models.Publisher{
		BookID: book.ID,
		Name:   faker.Word(),
	}
}
