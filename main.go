package main

import (
	"bank/repository"
	"bank/service"

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

	// testCreate(customerRepo)
	// testGetAll(customerRepo)
	// testGetById(customerRepo, 1)
	// testDelete(customerRepo, 2)

}

// func testCreate(customerRepo repository.CustomerRepository) {

// 	err := customerRepo.Create(&repository.Customer{
// 		CustomerID: 2,
// 		Name:       "sloth",
// 		Status:     1,
// 	})

// 	fmt.Println(err)

// }

// func testGetAll(customerRepo repository.CustomerRepository) {

// 	customers, err := customerRepo.GetAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(customers)

// }

// func testGetById(customerRepo repository.CustomerRepository, id int) {

// 	customer, err := customerRepo.GetById(id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(*customer)

// }

// func testDelete(customerRepo repository.CustomerRepository, id int) {

// 	err := customerRepo.Delete(id)

// 	fmt.Println(err)

// }
