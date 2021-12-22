package routes

import (
	"log"
	"net/http"

	"github.com/amrikmalhans/fanateapi/cmd/api/handlers"
	"github.com/gorilla/mux"
)

func Routes() {
    r := mux.NewRouter()
    r.HandleFunc("/hello", handlers.HandleHello).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}