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
	if err := db.AutoMigrate(&repository.Customer{}, &repository.Account{}); err != nil {
		panic(err)
	}

}
