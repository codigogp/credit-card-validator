package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateCard(t *testing.T) {
	tests := []struct {
		name            string
		method          string
		payload         string
		expectedStatus  int
		expectedValid   bool
		expectedNetwork string
	}{
		{
			name:            "Valid Visa Card",
			method:          http.MethodGet,
			payload:         `{"number":"4111111111111111"}`,
			expectedStatus:  http.StatusOK,
			expectedValid:   true,
			expectedNetwork: "Visa",
		},
		{
			name:            "Invalid Card Number",
			method:          http.MethodGet,
			payload:         `{"number":"1234567890123456"}`,
			expectedStatus:  http.StatusOK,
			expectedValid:   false,
			expectedNetwork: "Unknown",
		},
		{
			name:           "Invalid Method",
			method:         http.MethodPost,
			payload:        `{"number":"4111111111111111"}`,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodGet,
			payload:        `{"number":4111111111111111}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/validate", bytes.NewBufferString(tt.payload))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleValidateCreditCard)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response CreditCardResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Could not unmarshal response: %v", err)
				}

				if response.Valid != tt.expectedValid {
					t.Errorf("handler returned unexpected validity: got %v want %v",
						response.Valid, tt.expectedValid)
				}

				if response.Network != tt.expectedNetwork {
					t.Errorf("handler returned unexpected network: got %v want %v",
						response.Network, tt.expectedNetwork)
				}
			}
		})
	}
}
