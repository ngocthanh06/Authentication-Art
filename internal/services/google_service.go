package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ngocthanh06/authentication/internal/common"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type TokenInfoGoogle struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

var TokenInfoGoogleInstance *TokenInfoGoogle

const (
	UrlGoogleAuth        = "https://accounts.google.com/o/oauth2/v2/auth"
	UserInfoUrl          = "https://www.googleapis.com/oauth2/v1/userinfo"
	AccessTokenGoogleUrl = "https://oauth2.googleapis.com/token"
)

func GetGoogleOauthScope() []common.Scope {
	return []common.Scope{
		{"user_info_email", "https://www.googleapis.com/auth/userinfo.email"},
		{"user_info_profile", "https://www.googleapis.com/auth/userinfo.profile"},
		{"openid", "openid"},
	}
}

func GetGoogleUrlOauth2(oauthUrl string,
	clientId string,
	responseType string,
	state string,
	scope string,
	accessType string,
	redirectUri string,
	includeGrantedScopes string,
) string {
	return fmt.Sprintf("%s"+
		"?client_id=%s"+
		"&response_type=%s"+
		"&state=%s"+
		"&scope=%s"+
		"access_type=%s"+
		"&redirect_uri=%s"+
		"&prompt=consent"+
		"&include_granted_scopes=%s",
		oauthUrl,
		clientId,
		responseType,
		state,
		scope,
		accessType,
		redirectUri,
		includeGrantedScopes,
	)
}

type googleConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

var (
	once                 sync.Once
	googleConfigInstance *googleConfig
)

func InitGoogleConfig() *googleConfig {
	once.Do(func() {
		googleConfigInstance = &googleConfig{
			ClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_SECRET"),
			RedirectUri:  os.Getenv("GOOGLE_REDIRECT_URL"),
		}
	})

	return googleConfigInstance
}

/**
 * Get Token Google
 */
func GetTokensGoogle(code string) *TokenInfoGoogle {
	data := map[string]interface{}{
		"code":          code,
		"client_id":     googleConfigInstance.ClientId,
		"client_secret": googleConfigInstance.ClientSecret,
		"redirect_uri":  googleConfigInstance.RedirectUri,
		"grant_type":    "authorization_code",
	}

	jsonData, _ := json.Marshal(data)

	response, err := http.Post(AccessTokenGoogleUrl, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request: %v", err)

		return nil
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK || err != nil {
		fmt.Printf("Error reading response: %v", err)

		return nil
	}

	err = json.Unmarshal(body, &TokenInfoGoogleInstance)

	if err != nil {
		fmt.Printf("Error reading response: %v", err)

		return nil
	}

	return TokenInfoGoogleInstance
}

func GetGoogleUserInfo(accessToken string) []byte {
	resp, err := http.Get(UserInfoUrl + "?access_token=" + accessToken)

	if err != nil {
		fmt.Printf("error sending request: %v", err)

		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		fmt.Printf("failed to get user info: %d, body: %s", resp.StatusCode, body)
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error reading response body: %v", err)

		return nil
	}

	return body
}
