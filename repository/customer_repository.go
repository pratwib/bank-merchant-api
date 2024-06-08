package repository

import (
	"bank-merchant-api/model/domain"
	"github.com/jinzhu/gorm"
)

type CustomerRepository interface {
	Save(customer *domain.Customer) error
	FindByEmail(email string) (*domain.Customer, error)
	UpdateBalance(customerID string, amount float64) error
	BeginTransaction() (*gorm.DB, error)
}
