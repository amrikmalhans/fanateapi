package routes

import (
	"log"
	"net/http"

	"github.com/amrikmalhans/fanateapi/cmd/api/handlers"
	"github.com/gorilla/mux"
)

func Routes() {
    r := mux.NewRouter()
	r.HandleFunc("/", handlers.HandleIndex).Methods("GET")
    r.HandleFunc("/login", handlers.HandleLogin).Methods("GET")
	r.HandleFunc("/callback/", handlers.HandleCallback).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":9000", r))
}