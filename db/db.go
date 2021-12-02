package db

import (
	"main/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB() {
	dsn := "user=postgres password=12341234 dbname=miniecom port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.Demo{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Product{})
	// database.AutoMigrate(&model.Transaction{})

	db = database
}
