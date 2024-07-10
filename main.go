package main

import (
	"bank/repository"
	"bank/service"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", "__sqlite3.db")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	customerRepo, err := repository.NewCustomerRepository(db)
	if err != nil {
		panic(nil)
	}

	customerService := service.NewCustomerService(customerRepo)

	customerResponses, err := customerService.GetCustomers()
	if err != nil {
		fmt.Println("Error from query operation.")
		log.Fatal(err)
	}

	fmt.Println("Successfully queried.")
	log.Println(customerResponses)

}
