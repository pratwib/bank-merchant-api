package repository

import "bank-merchant-api/model/domain"

type MerchantRepository interface {
	Save(merchant *domain.Merchant) error
	FindByID(id string) (*domain.Merchant, error)
	UpdateBalance(merchantID string, amount float64) error
	Delete(id string) error
}
