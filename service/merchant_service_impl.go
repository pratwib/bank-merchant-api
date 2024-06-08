package service

import (
	"bank-merchant-api/model/domain"
	"bank-merchant-api/repository"
)

type MerchantServiceImpl struct {
	merchantRepository repository.MerchantRepository
	historyRepository  repository.HistoryRepository
}

func NewMerchantService(merchantRepository repository.MerchantRepository, historyRepository repository.HistoryRepository) MerchantService {
	return &MerchantServiceImpl{merchantRepository: merchantRepository, historyRepository: historyRepository}
}

func (service *MerchantServiceImpl) CreateMerchant(merchant *domain.Merchant) error {
	return service.merchantRepository.Save(merchant)
}

func (service *MerchantServiceImpl) GetMerchant(id string) (*domain.Merchant, error) {
	return service.merchantRepository.FindByID(id)
}

func (service *MerchantServiceImpl) DeleteMerchant(id string) error {
	return service.merchantRepository.Delete(id)
}
