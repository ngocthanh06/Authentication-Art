package models

import (
	"github.com/dgrijalva/jwt-go"
)

var Auth *User

type Credentials struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type ResponseDataLoginSuccess struct {
	Data        interface{} `json:"data,omitempty"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresIn   int         `json:"expires_in"`
}
