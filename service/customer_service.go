package service

import (
	"bank/logs"
	"bank/repository"
	"database/sql"
	"errors"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(custRepository repository.CustomerRepository) customerService {
	return customerService{
		customerRepo: custRepository,
	}
}

func (c customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := c.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerId: customer.CustomerID,
			Name:       customer.Name,
		}
		customerResponses = append(customerResponses, custResponse)
	}
	return customerResponses, nil
}
func (c customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := c.customerRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Customer not found.")
		}
		logs.Error(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerId: customer.CustomerID,
		Name:       customer.Name,
	}
	return &custResponse, nil
}
