package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var (
	ClientID     = ""
	ClientSecret = ""
)

func getToken(refreshToken string) (string, error) {
	var token struct {
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

	return token.AccessToken, nil
}

func genAuthString(user, refreshToken string) string {
	token, _ := getToken(refreshToken)
	auth := strings.Join([]string{"user=", user, "\x01auth=Bearer ", token, "\x01\x01"}, "")
	return auth
}
