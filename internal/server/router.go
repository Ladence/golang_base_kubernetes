package server

import (
	"encoding/json"
	"fmt"
	"github.com/Ladence/golang_base_kubernetes/internal/bl"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	response := bl.Home()
	body, err := json.Marshal(response)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, fmt.Sprintf("Could not encode info data: %v", err), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Error on message sending: %v", err)
	}
}