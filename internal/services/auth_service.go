package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/authentication/internal/config"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceInterface interface {
	Login(ctx *gin.Context) (string, error)
}

type AuthService struct {
	authRepository *repositories.DbStorage
	userRepository *repositories.DbStorage
}

func NewAuthService(authRepository *repositories.DbStorage) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		userRepository: authRepository,
	}
}

// Login /**
func (authService *AuthService) Login(ctx *gin.Context, data *models.Credentials) (*models.ResponseDataLoginSuccess, error) {
	// check exist user
	result, err := authService.userRepository.GetUserByConditions(ctx, map[string]interface{}{
		"email": data.Email,
	})

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
		TokenType:   config.TokenType,
		ExpiresIn:   minutesLeft,
	}

	return &success, nil
}
