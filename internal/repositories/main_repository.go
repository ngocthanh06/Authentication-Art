package repositories

import "gorm.io/gorm"

type DbStorage struct {
	db *gorm.DB
}
