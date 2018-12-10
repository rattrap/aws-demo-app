package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"status,omitempty"`
}

func GetHome(w http.ResponseWriter, _ *http.Request) {
	b := HomeResponse{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(b)
}

type StatusResponse struct {
	Status string `json:"status,omitempty"`
}

func GetStatus(w http.ResponseWriter, _ *http.Request) {
	b := StatusResponse{Status: "idle"}
	json.NewEncoder(w).Encode(b)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetHome).Methods("GET")
	router.HandleFunc("/status", GetStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":8123", router))
}
