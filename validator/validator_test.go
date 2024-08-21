package validator

import "testing"

func TestValidateLuhn(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected bool
	}{
		{"Valid Visa", "4111111111111111", true},
		{"Valid MasterCard", "5500000000000004", true},
		{"Invalid number", "1234567890123456", false},
		{"Empty string", "", false},
		{"Non-numeric string", "abcdefghijklmnop", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateLuhn(tt.number); got != tt.expected {
				t.Errorf("ValidateLuhn(%s) = %v, want %v", tt.number, got, tt.expected)
			}
		})
	}
}

func TestIdentifyCardType(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected string
	}{
		{"Visa", "4111111111111111", "Visa"},
		{"MasterCard", "5500000000000004", "MasterCard"},
		{"American Express", "371449635398431", "American Express"},
		{"Discover", "6011000990139424", "Discover"},
		{"JCB", "3530111333300000", "JCB"},
		{"UnionPay", "6200000000000005", "UnionPay"},
		{"RuPay", "6070000000000000", "RuPay"},
		{"Unknown", "9876543210123456", "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdentifyCardType(tt.number); got != tt.expected {
				t.Errorf("IdentifyCardType(%s) = %v, want %v", tt.number, got, tt.expected)
			}
		})
	}
}
