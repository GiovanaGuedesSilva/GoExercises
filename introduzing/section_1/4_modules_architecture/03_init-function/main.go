package main

/*
	Demonstração do uso da função init() em Go.

	- A função init() é executada automaticamente antes da função main().
	- Pode ser usada para inicializações (ex: carregar configs, verificar variáveis).
	- Pode haver múltiplas init() em diferentes arquivos e pacotes.
	- A ordem de execução segue a ordem de importação dos pacotes.
*/

import "fmt"

// init é chamada automaticamente antes do main
func init() {
	fmt.Println("Executando init()... preparando o ambiente")
}

func main() {
	fmt.Println("Executando main()... aplicação rodando")
}
