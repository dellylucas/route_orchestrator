package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", a.getProduct).Methods(http.MethodGet)
	r.HandleFunc("/dashboard-delivery/{id}", a.getDelivery).Methods(http.MethodGet)

}
