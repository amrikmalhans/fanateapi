package handlers

import (
	"net/http"

	"github.com/amrikmalhans/fanateapi/cmd/api/services"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	cred := services.Login()

	http.Redirect(w, r, 
		"https://accounts.spotify.com/authorize?client_id="+cred.Client_ID+"&response_type=code&redirect_uri="+cred.Redirect_URI+"&scope="+cred.Scope+"&state="+cred.State, http.StatusFound)

}