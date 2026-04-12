package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // ❗ yahan change

func Connect() {

	dsn := "host=localhost user=postgres password=root dbname=Khadbhandar port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		println("Database not connected, running without DB")
		return
	}

	DB = db
}
