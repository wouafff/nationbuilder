package nationbuilder

import (
	"fmt"
	"net/http"
)

type TokenResult struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	CreatedAt    int32  `json:"created_at"`
	RefreshToken string `json:"refresh_token"`
}
type TokenQuery struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	RefreshToken string `json:"refresh_token"`
}

// Retrieve a token by calling the dedicated method
func (n *Client) Token(grantType string, code string, clientId string, clientSecret string, redirectUri string) TokenResult {
	u := "/oauth/token"
	req := n.getRequest("POST", u, nil)
	token := TokenResult{}
	tokenQuery := TokenQuery{GrantType: grantType, Code: code, ClientId: clientId, ClientSecret: clientSecret, RedirectUri: redirectUri}
	result := n.create(tokenQuery, req, &token, http.StatusOK)
	if result.Err != nil {
		fmt.Println(result.Err)
	}
	return token
}

// Retrieve a token via refreshtoken by calling the dedicated method
func (n *Client) RefreshToken(grantType string, refreshToken string, clientId string, clientSecret string) TokenResult {
	u := "/oauth/token"
	req := n.getRequest("POST", u, nil)
	token := TokenResult{}
	tokenQuery := TokenQuery{GrantType: grantType, RefreshToken: refreshToken, ClientId: clientId, ClientSecret: clientSecret}
	fmt.Println(tokenQuery)
	result := n.create(tokenQuery, req, &token, http.StatusOK)
	fmt.Println(result)
	if result.Err != nil {
		fmt.Println(result.Err)
	}
	return token
}
