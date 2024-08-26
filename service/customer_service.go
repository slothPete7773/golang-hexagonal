package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
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
		return nil, errs.NewUnexpectedError()
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
			// return nil, errors.New("Customer not found.")
			return nil, errs.NewNotFoundError("Customer not found.")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerId: customer.CustomerID,
		Name:       customer.Name,
	}
	return &custResponse, nil
}
