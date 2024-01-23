package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type API struct {
	Data Products `json:"data"`
}

type Products struct {
	Request_id string        `json:"request_id"`
	Products   []DataProduct `json:"products"`
}

type DataProduct struct {
	Typed             string `json:"type"`
	Name              string `json:"name"`
	Number            string `json:"number"`
	Available_balance int    `json:"available_balance"`
	Status            string `json:"status"`
}

var account = DataProduct{"Account", "Cuenta de ahorros", "1234567", 300000, "1"}
var credit = DataProduct{"Credit Card", "Cuenta libre inversion", "3344551", 20000, "1"}

var product = Products{"72241424214123", []DataProduct{account, credit}}
var datas = API{product}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	_, err := strconv.Atoi(limitParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(datas)
}

func (a *API) getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datas)
	log.Println("Call")
}

func (a *API) getDelivery(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(datas)
}
