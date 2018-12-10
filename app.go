package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

type env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type home struct {
	Hostname string `json:"hostname,omitempty"`
	Envs     []env  `json:"envs,omitempty"`
}

type status struct {
	Status string `json:"status,omitempty"`
}

func getHome(w http.ResponseWriter, _ *http.Request) {
	hostname, _ := os.Hostname()

	envs := []env{}
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		envs = append(envs, env{Key: pair[0], Value: pair[1]})
	}

	b := home{Hostname: hostname, Envs: envs}
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
