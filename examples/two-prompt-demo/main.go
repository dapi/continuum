package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response{Service: "continuum-demo", Status: "ok"})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	log.Println("continuum demo listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
