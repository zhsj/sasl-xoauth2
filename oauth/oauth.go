package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func getToken(refreshToken string) (string, error) {
	var token struct {
		Error       string `json:"error"`
		ErrorDesc   string `json:"error_description"`
		AccessToken string `json:"access_token"`
	}
	values := url.Values{
		"client_id":     {ClientID},
		"client_secret": {ClientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	}
	resp, err := http.PostForm("https://www.googleapis.com/oauth2/v4/token", values)
	if err != nil {
		return "", err
	}

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	if token.Error != "" {
		return "", fmt.Errorf("refresh token failed: %s, %s", token.Error, token.ErrorDesc)
	}
	return token.AccessToken, nil
}

func PrintRefreshToken() {
	values := url.Values{
		"client_id":     {ClientID},
		"redirect_uri":  {"urn:ietf:wg:oauth:2.0:oob"},
		"scope":         {"https://mail.google.com/"},
		"response_type": {"code"},
	}
	authUrl, _ := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")
	authUrl.RawQuery = values.Encode()
	fmt.Println("To authorize token, visit this url and follow the directions:")
	fmt.Println(" ", authUrl)
	fmt.Print("Enter verification code: ")
	var code string
	fmt.Scan(&code)
	values = url.Values{
		"client_id":     {ClientID},
		"client_secret": {ClientSecret},
		"code":          {code},
		"redirect_uri":  {"urn:ietf:wg:oauth:2.0:oob"},
		"grant_type":    {"authorization_code"},
	}
	var token struct {
		Error        string `json:"error"`
		ErrorDesc    string `json:"error_description"`
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}
	resp, err := http.PostForm("https://www.googleapis.com/oauth2/v4/token", values)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		fmt.Println(err)
		return
	}
	if token.Error != "" {
		fmt.Println("Error: ", token.Error, token.ErrorDesc)
	} else {
		fmt.Println("Use this secret as password in Muttrc: ", token.RefreshToken)
		// fmt.Println("The temporary access token is: ", token.AccessToken)
	}
}

func GenAuthString(user, refreshToken string) string {
	token, err := getToken(refreshToken)
	if err != nil {
		return err.Error()
	}
	auth := strings.Join([]string{"user=", user, "\x01auth=Bearer ", token, "\x01\x01"}, "")
	return auth
}
