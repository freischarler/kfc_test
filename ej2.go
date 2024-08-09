package main

import (
	"fmt"
	"strings"
)

// Función para encriptar el mensaje
func encryptMessage(key, message string) string {
	if message == "" {
		return ""
	}

	if key == "" {
		key = "DCJ"
	}

	// Definir las vocales
	vowels := "AEIOUaeiou"

	// Crear un builder para la cadena resultante
	var result strings.Builder

	// Iterar sobre el mensaje
	for _, char := range message {
		// Convertir el char a string para comparación
		charStr := string(char)
		if strings.Contains(vowels, charStr) {
			result.WriteString(key)
		}
		result.WriteString(charStr)
	}

	return result.String()
}

func main() {
	// Ejemplo de uso
	key := "dcj"
	message := "I love prOgrAmming!"
	encryptedMessage := encryptMessage(key, message)
	fmt.Println(encryptedMessage)
}
