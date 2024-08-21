# ⚡ Credit Card Validator ⚡

## 🤔 ¿Qué es Credit Card Validator?

Credit Card Validator es una librería de código abierto en Go diseñada para simplificar la validación e identificación de tarjetas de crédito. Ofrece una API HTTP sencilla para validar números de tarjeta utilizando el algoritmo de Luhn e identificar la red de pago asociada.

## Características

* **Validación Luhn**: Implementa el algoritmo de Luhn para verificar la validez de los números de tarjeta.
* **Identificación de red**: Reconoce las principales redes de pago como Visa, MasterCard, American Express, y más.
* **API HTTP**: Proporciona un endpoint GET simple para validar tarjetas a través de solicitudes JSON.
* **Respuestas JSON**: Devuelve resultados en formato JSON para fácil integración con otros sistemas.
* **Pruebas unitarias**: Incluye un conjunto completo de pruebas para asegurar la fiabilidad del código.

## Instalación

```bash
go get github.com/codigogp/credit-card-validator
```

## Uso

### Iniciar el servidor

```go
package main

import (
    "log"
    "net/http"

    "github.com/codigogp/credit-card-validator/handler"
)

func main() {
    http.HandleFunc("/validate", handler.ValidateCard)
    log.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Realizar una solicitud

Usa curl o cualquier cliente HTTP para hacer una solicitud GET:

```bash
curl -X GET -H "Content-Type: application/json" -d '{"number":"4111111111111111"}' http://localhost:8080/validate
```

### Respuesta

```json
{
  "valid": true,
  "network": "Visa",
  "number": "4111111111111111"
}
```

## Estructura del Proyecto

```
credit-card-validator/
├── main.go
├── handler/
│   ├── handler.go
│   └── handler_test.go
├── validator/
│   ├── validator.go
│   └── validator_test.go
├── go.mod
└── README.md
```

## Contribución

¡Las contribuciones son bienvenidas! Por favor, lee las directrices de contribución antes de enviar un pull request.

1. Fork el proyecto
2. Crea tu rama de características (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Pruebas

Para ejecutar las pruebas:

```bash
go test ./...
```

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## Contacto

Link del Proyecto: [https://github.com/codigogp/credit-card-validator](https://github.com/codigogp/credit-card-validator)