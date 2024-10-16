package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ngocthanh06/authentication/internal/common"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/services"
	"log"
	"net/http"
)

type userHandler struct {
	userService services.UserServiceInterface
}

func UserHandler(userService services.UserServiceInterface) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (userService *userHandler) CreateUser(ctx *gin.Context) {
	var user models.UserCreation

	if err := ctx.ShouldBindWith(&user, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseValidationErrors(http.StatusBadRequest, err))
		return
	}

	result, err := userService.userService.UserCreate(&user)

	if err != nil {
		log.Print("Cannot create user: ", err.Error())

		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, err.Error()))

		return
	}

	ctx.JSON(http.StatusCreated, common.ResponseSuccessfully(map[string]interface{}{"users": result}, "created success"))
}

func (userService *userHandler) GetUsers(ctx *gin.Context) {
	results, err := userService.userService.UserList()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, err.Error()))

		return
	}

	ctx.JSON(http.StatusCreated, common.ResponseSuccessfully(map[string]interface{}{"users": results}, "User list"))
}

func (userService *userHandler) FindUser(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Println("id stirng", id)
	result, err := userService.userService.FindUser(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, "User not found!"))

		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccessfully(map[string]interface{}{
		"user": result,
	}, "Get user success"))
}
