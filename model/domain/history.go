package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type History struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	CustomerID string    `gorm:"type:varchar(100);not null"`
	Activity   string    `gorm:"type:varchar(100);not null"`
	Timestamp  string    `gorm:"type:timestamp;not null"`
}

func (history *History) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
