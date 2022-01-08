package seeders

import (
	"github.com/idriscahyono/grpc-bookstore/app/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.BookFaker(db)},
		{Seeder: fakers.PublisherFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, s := range RegisterSeders(db) {
		err := db.Debug().Create(s.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
