package repositories

import (
	"gorm.io/gorm"
)

func NewAuthRepository(dtb *gorm.DB) *DbStorage {
	return &DbStorage{
		db: dtb,
	}
}
