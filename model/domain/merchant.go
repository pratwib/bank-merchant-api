package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Merchant struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name    string    `gorm:"type:varchar(100);not null"`
	Balance float64   `gorm:"type:decimal(10,2)"`
}

func (merchant *Merchant) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
