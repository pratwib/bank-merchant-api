package repository

import (
	"bank-merchant-api/model/domain"
	"github.com/jinzhu/gorm"
	"time"
)

type HistoryRepositoryImpl struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &HistoryRepositoryImpl{db: db}
}

func (repo *HistoryRepositoryImpl) LogHistory(history *domain.History) error {
	history.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	return repo.db.Save(history).Error
}
