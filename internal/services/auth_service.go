package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ngocthanh06/authentication/internal/config"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/repositories"
	"github.com/ngocthanh06/authentication/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	authRepository *repositories.DbStorage
	userRepository repositories.UserRepositoryInterface
}

func NewAuthService(authRepository *repositories.DbStorage) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

// Login /**
func (authService *AuthService) Login(data *models.Credentials) (*models.ResponseDataLoginSuccess, error) {
	// check exist user
	result, err := authService.userRepository.GetUserByConditions(map[string]interface{}{
		"email": data.Email,
	})

	fmt.Print(result)

	if err != nil {
		fmt.Println("user err", err)
		return nil, err
	}

	// check has password
	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password)); err != nil {
		return nil, errors.New("Password or email is not correct!")
	}

	expirationTime := time.Now().Add(config.EnvKey.TTL * time.Minute)
	minutesLeft := int(time.Until(expirationTime).Minutes())

	claims := &models.Claims{
		Email: result.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Subject:   result.Id.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.EnvKey.JwtKey)

	if err != nil {
		fmt.Println("token string", err)

		return nil, err
	}

	success := models.ResponseDataLoginSuccess{
		AccessToken: tokenString,
		TokenType:   utils.TokenType,
		ExpiresIn:   minutesLeft,
	}

	return &success, nil
	return nil, nil
}
