package repositories

import (
	"fmt"
	"github.com/ngocthanh06/authentication/internal/models"
	"gorm.io/gorm"
)

type UserDbStorage struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	UserList(data map[string]interface{}) (*[]models.User, error)
	GetUserByConditions(conditions map[string]interface{}) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	FindUser(result map[string]interface{}) (*models.User, error)
}

func NewUserRepository(dtb *gorm.DB) UserRepositoryInterface {
	return &UserDbStorage{
		db: dtb,
	}
}

func (storage *UserDbStorage) UserList(data map[string]interface{}) (*[]models.User, error) {
	var users []models.User
	results := storage.db.Where(data).Find(&users)

	fmt.Println(results)

	return &users, nil
}

func (storage *UserDbStorage) GetUserByConditions(conditions map[string]interface{}) (*models.User, error) {
	var user models.User
	result := storage.db.Where(conditions).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (storage *UserDbStorage) CreateUser(user *models.User) (*models.User, error) {
	result := storage.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (storage *UserDbStorage) FindUser(result map[string]interface{}) (*models.User, error) {
	var user models.User

	results := storage.db.First(&user, result)

	if results.Error != nil {
		return nil, results.Error
	}

	return &user, nil
}
