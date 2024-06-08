package service

import "bank-merchant-api/model/domain"

type MerchantService interface {
	CreateMerchant(merchant *domain.Merchant) error
	GetMerchant(id string) (*domain.Merchant, error)
	DeleteMerchant(id string) error
}
