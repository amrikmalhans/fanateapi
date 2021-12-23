package services

import (
	"github.com/amrikmalhans/fanateapi/cmd/api/helpers"
)

const (
	CLIENT_ID     = ""
	REDIRECT_URI  = "http://localhost:9000/callback/"
	SCOPE 		  = "user-read-private user-read-email"
	CLIENT_SECRET = ""
)

type Credentials struct {
	Client_ID string `json:"client_id"`
	Redirect_URI string `json:"redirect_uri"`
	Scope string `json:"scope"`
	State string `json:"rand"`

} 

func Login() *Credentials {
	state, err := helpers.GenerateRandomString(16)

	if err != nil {
		panic(err)
	}

	 cred := &Credentials{
		Client_ID: CLIENT_ID,
		Redirect_URI: REDIRECT_URI,
		Scope: SCOPE,
		State: state,
	}

	
	return cred
}