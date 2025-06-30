package main

/*
	Exemplo de uso de um package chamado mathutil.

	- Importamos o pacote local usando o nome do módulo.
	- Chamamos as funções exportadas (letra maiúscula).
*/

import (
	"fmt"

	"example.com/modularapp/mathutil"
)

func main() {
	soma := mathutil.Somar(3, 4)
	sub := mathutil.Subtrair(10, 5)

	fmt.Println("Resultado da soma:", soma)
	fmt.Println("Resultado da subtração:", sub)
}
