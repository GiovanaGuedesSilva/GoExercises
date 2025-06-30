package greetings

import "fmt"

// Hello retorna uma saudação personalizada
func Hello(name string) string {
	return fmt.Sprintf("Olá, %s! Bem-vinda ao projeto modular!", name)
}
