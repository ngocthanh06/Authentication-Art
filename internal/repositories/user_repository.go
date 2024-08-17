package repositories

import (
	"context"
	"fmt"
	"github.com/ngocthanh06/authentication/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserCreate(ctx context.Context, data *models.TodoItemCreation) error
	UserList(ctx context.Context, data *models.TodoItemCreation) error
	CreateUser(ctx context.Context, data models.UserCreation) (*models.User, error)
	FindUser(ctx *context.Context, data *models.User) (*models.User, error)
}

func NewUserRepository(dtb *gorm.DB) *DbStorage {
	return &DbStorage{
		db: dtb,
	}
}

func (storage *DbStorage) UserList(ctx context.Context, data map[string]interface{}) (*[]models.User, error) {
	var users []models.User
	results := storage.db.Where(data).Find(&users)

	fmt.Println(results)

	return &users, nil
}

func (storage *DbStorage) GetUserByConditions(ctx context.Context, conditions map[string]interface{}) (*models.User, error) {
	var user models.User
	result := storage.db.Where(conditions).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (storage *DbStorage) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	result := storage.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (storage *DbStorage) FindUser(ctx context.Context, result map[string]interface{}) (*models.User, error) {
	var user models.User

	results := storage.db.First(&user, result)

	if results.Error != nil {
		return nil, results.Error
	}

	return &user, nil
}
