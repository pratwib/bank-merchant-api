package service

import "bank-merchant-api/model/domain"

type CustomerService interface {
	Register(customer *domain.Customer) error
	Login(email, password string) (string, error)
	AddBalance(customerID string, amount float64) error
	ProcessPayment(customerID, merchantID string, amount float64) error
}
