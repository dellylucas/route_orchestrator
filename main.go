package main

import (
	"dellyscotia/api"
	"net/http"

	"encoding/json"
	"log"

	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"Helo W\"}")
}

func main() {
	r := mux.NewRouter()

	a := &api.API{}
	a.RegisterRoutes(r)

	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	log.Println("listener...")
	srv.ListenAndServe()
}
