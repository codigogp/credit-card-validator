package validator

import "strings"

// ValidateLuhn verifica si un número de tarjeta de crédito es válido según el algoritmo de Luhn
func ValidateLuhn(cardNumber string) bool {

	// Primero, verificamos si la cadena está vacía
	if len(cardNumber) == 0 {
		return false
	}

	var sum int
	parity := len(cardNumber) % 2

	for i, digit := range cardNumber {
		if digit < '0' || digit > '9' {
			return false // Caracter no válido
		}

		d := int(digit - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}

// IdentifyCardType determina el tipo de tarjeta basado en su número
func IdentifyCardType(number string) string {
	number = strings.TrimSpace(number)

	switch {
	case len(number) == 15 && (strings.HasPrefix(number, "34") || strings.HasPrefix(number, "37")):
		return "American Express"
	case len(number) == 16 && strings.HasPrefix(number, "4"):
		return "Visa"
	case len(number) >= 16 && len(number) <= 19 && strings.HasPrefix(number, "5") && number[1] >= '1' && number[1] <= '5':
		return "MasterCard"
	case len(number) == 16 && (strings.HasPrefix(number, "6011") || strings.HasPrefix(number, "65")):
		return "Discover"
	case len(number) == 16 && (strings.HasPrefix(number, "2131") || strings.HasPrefix(number, "1800") || strings.HasPrefix(number, "35")):
		return "JCB"
	case len(number) >= 16 && len(number) <= 19 && strings.HasPrefix(number, "62"):
		return "UnionPay"
	case len(number) == 16 && strings.HasPrefix(number, "60"):
		return "RuPay"
	case len(number) == 16 && strings.HasPrefix(number, "9792"):
		return "Troy"
	default:
		return "Unknown"
	}
}
