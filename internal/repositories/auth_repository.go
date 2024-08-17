package repositories

import (
	"gorm.io/gorm"
)

type AuthRepository interface {
}

func NewAuthRepository(dtb *gorm.DB) *DbStorage {
	return &DbStorage{
		db: dtb,
	}
}
