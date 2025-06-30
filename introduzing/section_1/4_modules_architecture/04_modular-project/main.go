package main

/*
	Demonstração de projeto modular em Go com múltiplos pacotes.

	- Importação de pacotes internos (greetings e utils).
	- Separação por responsabilidade.
*/

import (
	"fmt"
	"modular-project/greetings"
	"modular-project/utils"
)

func main() {
	msg := greetings.Hello("Giovana")
	fmt.Println(msg)

	soma := utils.Add(5, 7)
	fmt.Println("Soma:", soma)
}
