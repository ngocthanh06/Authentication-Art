package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/authentication/internal/common"
	"github.com/ngocthanh06/authentication/internal/services"
	"net/http"
	"sync"
)

var once sync.Once

var accessType = "offline"
var includeGrantedScopes = "true"
var responseType = "code"
var state = "state"

var scopes []common.Scope

type providerConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

type oauthUrl2 struct {
	providerConfig       providerConfig
	oauthUrl             string
	responseType         string
	state                string
	scope                string
	accessType           string
	includeGrantedScopes string
}

var oauthUrl oauthUrl2
var providerConfigInstance *providerConfig

func InitProviderConfig(provider string) {
	switch provider {
	case "google":
		once.Do(func() {
			providerConfigInstance = (*providerConfig)(services.InitGoogleConfig())
			scopes = services.GetGoogleOauthScope()
			oauthUrl = oauthUrl2{
				oauthUrl:             services.UrlGoogleAuth,
				responseType:         responseType,
				state:                state,
				scope:                compareScope(),
				accessType:           accessType,
				includeGrantedScopes: includeGrantedScopes,
				providerConfig:       *providerConfigInstance,
			}
		})
		break
	case "github":
		break
	case "facebook":
		break
	}
}

func compareScope() string {
	str := ""

	for _, val := range scopes {
		str += val.Value + "&"
	}

	return str
}

func getUrlOauth2(provider string) string {
	var urlOauth string
	switch provider {
	case "google":
		urlOauth = services.GetGoogleUrlOauth2(oauthUrl.oauthUrl,
			oauthUrl.providerConfig.ClientId,
			oauthUrl.responseType,
			oauthUrl.state,
			oauthUrl.scope,
			oauthUrl.accessType,
			oauthUrl.providerConfig.RedirectUri,
			oauthUrl.includeGrantedScopes)
		break
	case "github":
		break
	case "facebook":
		break
	default:

		break
	}

	return urlOauth
}

func getUserInfoProvider(provider string, code string) []byte {
	var userInfo []byte

	switch provider {
	case "google":
		tokenInstance := services.GetTokensGoogle(code)

		if tokenInstance == nil {
			return nil
		}

		userInfo = services.GetGoogleUserInfo(tokenInstance.AccessToken)
		break
	case "github":
		break
	case "facebook":
		break
	default:
		break
	}

	return userInfo
}

func CallbackProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")
	InitProviderConfig(provider)
	code := ctx.Query("code")

	userInfo := getUserInfoProvider(provider, code)

	if userInfo == nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, errors.New("cannot get user info"), "cannot get user info"))

		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccessfully(string(userInfo), "get user info success"))

	return
}

func RedirectProviderLogin(ctx *gin.Context) {
	provider := ctx.Param("provider")
	InitProviderConfig(provider)

	var urlOauth2 = getUrlOauth2(provider)

	if urlOauth2 == "" {
		ctx.JSON(http.StatusBadRequest, common.ResponseError(http.StatusBadRequest, errors.New("cannot get url oauth"), "cannot get url oauth"))

		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccessfully(urlOauth2, "Redirect success"))
	return
}
