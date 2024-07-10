package main

import (
	"bank/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Migrate() {
	db, err := gorm.Open(sqlite.Open("__sqlite3.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&repository.Customer{})

}
