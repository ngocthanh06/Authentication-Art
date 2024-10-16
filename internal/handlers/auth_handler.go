package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ngocthanh06/authentication/internal/common"
	"github.com/ngocthanh06/authentication/internal/models"
	"github.com/ngocthanh06/authentication/internal/providers"
	"net/http"
)

func Login(ctx *gin.Context) {
	var reqParams models.Credentials

	if err := ctx.ShouldBindWith(&reqParams, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseValidationErrors(http.StatusBadRequest, err))

		return
	}

	result, err := providers.AuthServ.Login(&reqParams)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, err, err.Error()))

		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccessfully(result, "Login success"))
}
