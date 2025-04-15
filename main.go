package main

import (
	"log"
	"net/http"
	"receipt-processor-challenge/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/receipts/process", routes.CreateReceipt).Methods("POST")

	r.HandleFunc("/receipts/{id}/points", routes.GetPoints).Methods("GET")

	// Run server on port 8080
	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
