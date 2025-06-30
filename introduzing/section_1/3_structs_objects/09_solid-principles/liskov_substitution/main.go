package main

import "fmt"

// Interface comum
type Notificador interface {
	Notificar(msg string)
}

type EmailNotificador struct{}

func (e EmailNotificador) Notificar(msg string) {
	fmt.Println("Notificando por Email:", msg)
}

type SmsNotificador struct{}

func (s SmsNotificador) Notificar(msg string) {
	fmt.Println("Notificando por SMS:", msg)
}

// Ambos podem ser usados no lugar da interface
func acionarNotificacao(n Notificador) {
	n.Notificar("Promoção disponível!")
}

func main() {
	acionarNotificacao(EmailNotificador{})
	acionarNotificacao(SmsNotificador{})
}
