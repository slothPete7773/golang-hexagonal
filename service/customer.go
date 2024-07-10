package service

type CustomerResponse struct {
	CustomerId int    `json:"customer_id"`
	Name       string `json:"customer_name"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
