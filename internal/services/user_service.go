package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepository repositories.UserRepositoryInterface
}

type UserServiceInterface interface {
	UserList() (*[]models.User, error)
	UserCreate(params *models.UserCreation) (*models.User, error)
	FindUser(id string) (*models.User, error)
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (userSer *userService) UserList() (*[]models.User, error) {
	result, err := userSer.userRepository.UserList(map[string]interface{}{})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (userService *userService) UserCreate(params *models.UserCreation) (*models.User, error) {
	// check user exists if exists error
	result, err := userService.userRepository.GetUserByConditions(map[string]interface{}{
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
	result, err = userService.userRepository.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (userSer *userService) FindUser(id string) (*models.User, error) {
	result, err := userSer.userRepository.FindUser(map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
