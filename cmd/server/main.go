package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ForgeFlow API is running"))
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Starting ForgeFlow server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}