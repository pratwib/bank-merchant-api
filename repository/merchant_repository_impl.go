package repository

import (
	"bank-merchant-api/model/domain"
	"github.com/jinzhu/gorm"
)

type MerchantRepositoryImpl struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &MerchantRepositoryImpl{db: db}
}

func (repository *MerchantRepositoryImpl) Save(merchant *domain.Merchant) error {
	return repository.db.Save(merchant).Error
}

func (repository *MerchantRepositoryImpl) FindByID(id string) (*domain.Merchant, error) {
	var merchant domain.Merchant
	if err := repository.db.Where("id = ?", id).First(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (repository *MerchantRepositoryImpl) UpdateBalance(merchantID string, amount float64) error {
	return repository.db.Model(&domain.Merchant{}).Where("id = ?", merchantID).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (repository *MerchantRepositoryImpl) Delete(id string) error {
	return repository.db.Where("id = ?", id).Delete(&domain.Merchant{}).Error
}
