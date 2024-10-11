package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/authentication/internal/common"
	"github.com/ngocthanh06/authentication/internal/config"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/providers"
	"net/http"
	"strings"
	"time"
)

var claims models.Claims

func checkAuthenticationHeader(c *gin.Context) error {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" {
		return errors.New("Authorization token is required")
	}

	reqToken := strings.Split(bearerToken, " ")[1]

	token, err := jwt.ParseWithClaims(reqToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.EnvKey.JwtKey, nil
	})

	// check timeout
	expiredTime := time.Until(time.Unix(claims.ExpiresAt, 0))

	if err != nil || !token.Valid || expiredTime.Seconds() <= 0 {
		return errors.New("Invalid or expired token")
	}

	return nil
}

func Recovery() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("start recovery")
		defer func() {
			if err := recover(); err != nil {
				if r := err.(error); r != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ResponseError(http.StatusInternalServerError,
						r,
						r.Error()),
					)
				}
			}
		}()
		fmt.Println("end recovery")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// recovery
		Recovery()

		// check authentication token header
		if err := checkAuthenticationHeader(c); err != nil {
			c.JSON(http.StatusUnauthorized, common.ResponseError(http.StatusUnauthorized, err, err.Error()))

			c.Abort()
			return
		}

		// check user exists
		user, err := providers.UserServ.FindUser(c, string(claims.Subject))

		if err != nil {
			c.JSON(http.StatusUnauthorized, common.ResponseError(http.StatusUnauthorized, err, err.Error()))

			c.Abort()
			return
		}

		// Set Auth global
		models.Auth = user

		fmt.Println("middleware user", user)

		c.Next()
	}
}
