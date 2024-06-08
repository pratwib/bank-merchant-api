package repository

import "bank-merchant-api/model/domain"

type HistoryRepository interface {
	LogHistory(history *domain.History) error
}
