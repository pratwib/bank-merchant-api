package service

import (
	"bank-merchant-api/model/domain"
	"bank-merchant-api/repository"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CustomerServiceImpl struct {
	customerRepository repository.CustomerRepository
	historyRepository  repository.HistoryRepository
	jwtSecret          string
}

func NewCustomerService(customerRepository repository.CustomerRepository, historyRepository repository.HistoryRepository, jwtSecret string) CustomerService {
	return &CustomerServiceImpl{customerRepository: customerRepository, historyRepository: historyRepository, jwtSecret: jwtSecret}
}

func (service *CustomerServiceImpl) Register(customer *domain.Customer) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	customer.Password = string(hashedPassword)

	if err := service.customerRepository.Save(customer); err != nil {
		return err
	}

	history := &domain.History{
		CustomerID: customer.ID.String(),
		Activity:   "register",
	}
	return service.historyRepository.LogHistory(history)
}

func (service *CustomerServiceImpl) Login(email, password string) (string, error) {
	customer, err := service.customerRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	token, err := service.generateJWT(customer.ID.String())
	if err != nil {
		return "", err
	}

	history := &domain.History{
		CustomerID: customer.ID.String(),
		Activity:   "login",
	}
	if err := service.historyRepository.LogHistory(history); err != nil {
		return "", err
	}

	return token, nil
}

func (service *CustomerServiceImpl) AddBalance(customerID string, amount float64) error {
	if err := service.customerRepository.UpdateBalance(customerID, amount); err != nil {
		return err
	}

	history := &domain.History{
		CustomerID: customerID,
		Activity:   fmt.Sprintf("add balance: %.2f", amount),
	}
	return service.historyRepository.LogHistory(history)
}

func (service *CustomerServiceImpl) ProcessPayment(customerID, merchantID string, amount float64) error {
	tx, err := service.customerRepository.BeginTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	customer := &domain.Customer{}
	merchant := &domain.Merchant{}

	if err := tx.Model(customer).Where("id = ?", customerID).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		return err
	}
	if err := tx.Model(merchant).Where("id = ?", merchantID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return err
	}

	merchantName := merchant.Name

	history := &domain.History{
		CustomerID: customerID,
		Activity:   fmt.Sprintf("payment to %s: %.2f", merchantName, amount),
	}
	if err := service.historyRepository.LogHistory(history); err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (service *CustomerServiceImpl) generateJWT(customerID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerID": customerID,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(service.jwtSecret))
}
