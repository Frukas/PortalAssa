package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//TokenHandler stores the authorization Token
type TokenHandler struct {
	Token     string
	NextToken string
	code      string
}

//Token is a struct of the json response of the token request
type Token struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
	Refreshtoken string `json:"refresh_token"`
}

//GetTokenFromCode retrive a token from a specific code using microsoft api
func (t *TokenHandler) GetTokenFromCode() {
	url := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	var payload Payload
	var token Token

	payload.Add("client_id", ClientID)
	payload.Add("scope", Scope)
	payload.Add("redirect_uri", RedirectURI)
	payload.Add("grant_type", "authorization_code")
	payload.Add("client_secret", ClientSecret)
	payload.Add("code", t.code)

	client := &http.Client{}
	tokenRequest, err := http.NewRequest("POST", url, payload.GetPayload())
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in the tokenRequest: %s", err))
	}

	tokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	tokenRequest.Header.Add("Cookie", "x-ms-gateway-slice=prod; stsservicecookie=ests; fpc=An9HL7q_iFBOqHpl5kLk6uRMKhPrAQAAADxgW9cOAAAA")

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(tokenRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &token); err != nil {
		log.Fatal(err)
	}

	t.Token = token.AccessToken
	t.NextToken = token.Refreshtoken

	// fmt.Println("Chegou no primeiro token")
	// fmt.Println(t.Token)
	//	fmt.Println(t.NextToken)

	go t.AutoRefreshToken()

}

//GetRefreshToken refreshes the token
func (t *TokenHandler) GetRefreshToken() {
	url := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	var payload Payload
	var token Token

	payload.Add("client_id", ClientID)
	payload.Add("redirect_uri", RedirectURI)
	payload.Add("grant_type", "refresh_token")
	payload.Add("client_secret", ClientSecret)
	payload.Add("refresh_token", t.NextToken)

	client := &http.Client{}
	tokenRequest, err := http.NewRequest("POST", url, payload.GetPayload())
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in the tokenRequest: %s", err))
	}

	tokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	tokenRequest.Header.Add("Cookie", "x-ms-gateway-slice=prod; stsservicecookie=ests; fpc=An9HL7q_iFBOqHpl5kLk6uRMKhPrAQAAADxgW9cOAAAA")

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(tokenRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &token); err != nil {
		log.Fatal(err)
	}

	t.Token = token.AccessToken
	t.NextToken = token.Refreshtoken
}

//AutoRefreshToken is responsible to call the function responsible for refresh the token ever 50 min.
func (t *TokenHandler) AutoRefreshToken() {
	for {
		time.Sleep(50 * time.Minute)
		t.GetRefreshToken()
	}
}
