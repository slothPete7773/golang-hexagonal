package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

// func (c *CustomerRepository) Create(customer *Customer) error {
// 	return nil
// }

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (c customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}

	query := "SELECT * FROM customers"
	err := c.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
func (c customerRepositoryDB) GetById(customerId int) (*Customer, error) {
	var customer Customer
	query := "SELECT * FROM customers WHERE customer_id = ?"
	err := c.db.Get(&customer, query, customerId)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
func (c customerRepositoryDB) Create(customer *Customer) error {
	query := "INSERT INTO customers (name, status) VALUES (:name, :status)"
	_, err := c.db.NamedExec(query, customer)
	if err != nil {
		return err
	}
	return nil
}
func (c customerRepositoryDB) Update(customer *Customer) error {
	query := "UPDATE customers SET name = :name, status = :status WHERE customer_id = :customer_id"
	_, err := c.db.NamedExec(query, customer)
	if err != nil {
		return err
	}
	return nil
}
func (c customerRepositoryDB) Delete(customerID int) error {
	query := "DELETE FROM customers WHERE customer_id = ?"
	_, err := c.db.Exec(query, customerID)
	if err != nil {
		return err
	}
	return nil
}
