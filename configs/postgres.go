package configs

import (
	"finalproject1/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Database *gorm.DB
}

func NewConnection(url string) *DbConnection {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}

	db.AutoMigrate(&model.Todo{})

	return &DbConnection{db}
}
