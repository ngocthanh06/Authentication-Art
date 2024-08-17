package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/authentication/internal/handlers"
	"github.com/ngocthanh06/authentication/internal/middleware"
	"net/http"
)

func CreateRoutes() {
	route := gin.Default()

	//route.Use(middleware.RequestTimeLogger())

	v1 := route.Group("v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK,
				gin.H{
					"message": "ping 1",
				},
			)

			return
		})

		v1.GET("/users", middleware.AuthMiddleware(), handlers.GetUsers)
		v1.GET("/user/:id", middleware.AuthMiddleware(), handlers.FindUser)
		v1.POST("/user", middleware.AuthMiddleware(), handlers.CreateUser)

		// Un-login
		v1.POST("/login", handlers.Login)
		v1.POST("/register", handlers.CreateUser)
	}

	route.Run(fmt.Sprintf(":%s", "8080"))
}
