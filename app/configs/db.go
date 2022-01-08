package configs

import (
	"fmt"
	"github.com/idriscahyono/grpc-bookstore/app/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	User   string
	Passwd string
	Host   string
	Port   int
	DBName string
}

var DB *gorm.DB
var err error

func InitDB() {
	db := config{
		User:   viper.Get("DB_USERNAME").(string),
		Passwd: viper.Get("DB_PASSWORD").(string),
		Host:   viper.Get("DB_HOST").(string),
		Port:   viper.Get("DB_PORT").(int),
		DBName: viper.Get("DB_NAME").(string),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User, db.Passwd, db.Host, db.Port, db.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	//Migration
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Publisher{})

	//Seeders
	//seeders.DBSeed(DB)
}
