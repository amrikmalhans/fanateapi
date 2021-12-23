package services

import (
	"github.com/amrikmalhans/fanateapi/cmd/api/helpers"
)

const (
	CLIENT_ID     = "013b8cf33386425d9835b0f8ed768125"
	REDIRECT_URI  = "http://localhost:9000/callback"
)

func Login() (string, error) {
	rand, err := helpers.GenerateRandomString(16)

	if err != nil {
		return "", err
	}

	return rand, err

}