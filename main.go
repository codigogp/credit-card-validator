package main

import (
	"log"
	"net/http"

	"github.com/codigogp/credit-card-validator/handler"
)

func main() {
	http.HandleFunc("/validate", handler.HandleValidateCreditCard)
	log.Println("Server starting on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
