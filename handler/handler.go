package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codigogp/credit-card-validator/validator"
)

type CreditCardRequest struct {
	Number string `json:"number"`
}

type CreditCardResponse struct {
	Valid   bool   `json:"valid"`
	Network string `json:"network"`
	Number  string `json:"number"`
}

func HandleValidateCreditCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el payload JSON
	var req CreditCardRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validar que el número de tarjeta no esté vacío
	if req.Number == "" {
		http.Error(w, "Card number is required", http.StatusBadRequest)
		return
	}

	// Ejecutar el algoritmo de Luhn y identificar la red de pago
	valid := validator.ValidateLuhn(req.Number)
	network := validator.IdentifyCardType(req.Number)

	// Preparar la respuesta
	resp := CreditCardResponse{
		Valid:   valid,
		Network: network,
		Number:  req.Number,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
