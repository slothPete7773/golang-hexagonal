package main

import (
	apihandler "bank/api_handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	customerHandler := apihandler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	fmt.Println("Running at localhost:8081")
	http.ListenAndServe(":8081", router)
	// customerResponses, err := customerService.GetCustomers()
	// if err != nil {
	// 	fmt.Println("Error from query operation.")
	// 	log.Fatal(err)
	// }

	// fmt.Println("Successfully queried.")
	// log.Println(customerResponses)

}
