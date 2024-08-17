package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ngocthanh06/authentication/internal/common"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/providers"
	"log"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user models.UserCreation

	if err := ctx.ShouldBindWith(&user, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseValidationErrors(http.StatusBadRequest, err))
		return
	}

	result, err := providers.UserServ.UserCreate(ctx, &user)

	if err != nil {
		log.Print("Cannot create user: ", err.Error())

		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, err.Error()))

		return
	}

	ctx.JSON(http.StatusCreated, common.ResponseSuccessfully(map[string]interface{}{"users": result}, "created success"))
}

func GetUsers(ctx *gin.Context) {
	results, err := providers.UserServ.UserList(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, err.Error()))

		return
	}

	ctx.JSON(http.StatusCreated, common.ResponseSuccessfully(map[string]interface{}{"users": results}, "User list"))
}

func FindUser(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Println("id stirng", id)
	result, err := providers.UserServ.FindUser(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, "User not found!"))

		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccessfully(map[string]interface{}{
		"user": result,
	}, "Get user success"))
}
