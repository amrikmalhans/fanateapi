package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AuthResp struct {
	Access_Token string `json:"access_token"`
	Token_Type string `json:"token_type"`
	Expires_In int `json:"expires_in"`
	Scope string `json:"scope"`
	Refresh_Token string `json:"refresh_token"`
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	data := url.Values{}
    data.Set("code", code)
    data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:9000/callback/")


	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode())) // URL-encoded payload
    if err != nil {
        log.Fatal(err)
    }

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Add("Authorization", "Basic ")


	res, err := client.Do(r)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatal(err)
    }

	var authResp AuthResp
	jsonErr := json.Unmarshal(body, &authResp)

	if jsonErr != nil {
		panic(jsonErr)
	}
	
}
