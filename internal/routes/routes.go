package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/authentication/internal/handlers"
	"github.com/ngocthanh06/authentication/internal/middleware"
	"github.com/ngocthanh06/authentication/internal/providers"
	"net/http"
)

func CreateRoutes() {
	route := gin.Default()

	// load controller
	userController := handlers.UserHandler(providers.UserServ)

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

		v1.GET("/users", middleware.AuthMiddleware(), userController.GetUsers)
		v1.GET("/user/:id", middleware.AuthMiddleware(), userController.FindUser)
		v1.POST("/user", middleware.AuthMiddleware(), userController.CreateUser)

		// Un-login
		v1.POST("/login", handlers.Login)
		v1.POST("/register", userController.CreateUser)

		// connect to social
		v1.GET("redirect/:provider", handlers.RedirectProviderLogin)
		v1.GET("auth/:provider/callback", handlers.CallbackProvider)
	}

	route.Run(fmt.Sprintf(":%s", "8080"))
}
