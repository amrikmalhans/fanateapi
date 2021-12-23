package handlers

import (
	"net/http"

	"github.com/amrikmalhans/fanateapi/cmd/api/services"
)


func HandleIndex(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code != "" && state != "" {
		services.SpotifyAuth(code, state)
	}

}