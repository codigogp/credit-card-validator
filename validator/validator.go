package validator

// ValidateLuhn verifica si un número de tarjeta de crédito es válido según el algoritmo de Luhn
func ValidateLuhn(cardNumber string) bool {

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
	switch {

	case len(number) == 15 && (number[:2] == "34" || number[:2] == "37"):
		return "American Express"

	case len(number) == 16 && number[:1] == "4":
		return "Visa"

	case len(number) >= 16 && len(number) <= 19 && number[:2] >= "51" && number[:2] <= "55":
		return "MasterCard"
	default:
		return "Unknown"
	}
}
