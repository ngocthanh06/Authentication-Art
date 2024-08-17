package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	UserCreate(ctx context.Context, data *models.TodoItemCreation) error
	UserList(ctx context.Context, data *models.TodoItemCreation) error
	FindUser(ctx *context.Context, data *models.User) (*models.User, error)
}

type UserService struct {
	userRepository *repositories.DbStorage
}

func NewUserService(repo *repositories.DbStorage) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (userSer *UserService) UserList(ctx context.Context) (*[]models.User, error) {
	result, err := userSer.userRepository.UserList(ctx, map[string]interface{}{})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (userService *UserService) UserCreate(ctx context.Context, params *models.UserCreation) (*models.User, error) {
	// check user exists if exists error
	result, err := userService.userRepository.GetUserByConditions(ctx, map[string]interface{}{
		"email": params.Email,
	})

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	if err == nil && result != nil {
		return nil, errors.New("Users exists !")
	}

	var hasPass []byte

	if params.Password != "" {
		hasPass, err = bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}
	}

	userId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	fmt.Println("userId", userId)

	user := models.User{
		Id:        userId,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Password:  string(hasPass),
		Email:     params.Email,
		Birthday:  params.Birthday,
	}

	// create new user
	result, err = userService.userRepository.CreateUser(ctx, &user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (userService *UserService) FindUser(ctx context.Context, id string) (*models.User, error) {
	result, err := userService.userRepository.FindUser(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
