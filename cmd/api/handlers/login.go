package handlers

import (
	"net/http"

	"github.com/amrikmalhans/fanateapi/cmd/api/services"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	rand, err := services.Login()
	
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(rand))

}