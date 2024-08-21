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
	Valid bool   `json:"valid"`
	Type  string `json:"type, omitempty"`
}

func HandleValidateCreditCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreditCardRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	valid := validator.ValidateLuhn(req.Number)
	cardType := validator.IdentifyCardType(req.Number)

	resp := CreditCardResponse{
		Valid: valid,
		Type:  cardType,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
