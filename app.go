package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type home struct {
	Message string `json:"status,omitempty"`
}

type status struct {
	Status string `json:"status,omitempty"`
}

func getHome(w http.ResponseWriter, _ *http.Request) {
	b := home{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(b)
}

func getStatus(w http.ResponseWriter, _ *http.Request) {
	b := status{Status: "idle"}
	json.NewEncoder(w).Encode(b)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/status", getStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":8123", router))
}
