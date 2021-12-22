package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	w.Write([]byte("Hello, " + name))

}