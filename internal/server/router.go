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
	router := mux.NewRouter()
	router.HandleFunc("/home", home).Methods("GET")

	router.HandleFunc("/healthz", healthCheck)
	router.HandleFunc("/readyz", readyCheck)

	return router
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

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readyCheck(w http.ResponseWriter, r *http.Request) {
	// here could be cache warming, resource preparation or smth like that
	w.WriteHeader(http.StatusOK)
}
