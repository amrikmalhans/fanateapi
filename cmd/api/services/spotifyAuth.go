package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)




func SpotifyAuth(code string, state string) {

	values := map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  "http://localhost:9000/",
	}

	jsonValue, _ := json.Marshal(values)

	req, err := http.Post("https://accounts.spotify.com/api/token", "application/x-www-form-urlencoded",
		bytes.NewBuffer(jsonValue),
	)

	if err != nil {
		panic(err)
	}


	req.Header.Set( "Authorization", "Basic "+CLIENT_ID+":"+CLIENT_SECRET)

	

}