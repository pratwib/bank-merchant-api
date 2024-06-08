package repository

import (
	"bank-merchant-api/model/domain"
	"github.com/jinzhu/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db: db}
}

func (repository *CustomerRepositoryImpl) Save(customer *domain.Customer) error {
	return repository.db.Save(customer).Error
}

func (repository *CustomerRepositoryImpl) FindByID(id string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := repository.db.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (repository *CustomerRepositoryImpl) FindByEmail(email string) (*domain.Customer, error) {
	var customer domain.Customer
	if err := repository.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (repository *CustomerRepositoryImpl) UpdateBalance(customerID string, amount float64) error {
	return repository.db.Model(&domain.Customer{}).Where("id = ?", customerID).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (repository *CustomerRepositoryImpl) BeginTransaction() (*gorm.DB, error) {
	tx := repository.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}
