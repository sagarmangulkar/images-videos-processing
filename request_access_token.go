package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiKey = "MyKey"
)

type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func main() {
	accessToken, err := requestAccessToken()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("AccessToken = ", accessToken)
}
func requestAccessToken() (string, error) {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	formData := strings.NewReader(form.Encode())
	url := fmt.Sprintf("https://api.clarifai.com/v2/models/" + apiKey + "/outputs")
	fmt.Printf("Url with my API key: %s\n", url)
	req, err := http.NewRequest("POST", url, formData)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Key MyKey")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var record TokenResp
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		return "", err
	}
	fmt.Printf("record: %s\n", record)
	return record.AccessToken, nil
}
