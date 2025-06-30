package main

/*
	Demonstração do uso de imports em Go.

	- Importação de pacotes padrão como "fmt", "math", "time".
	- Importação com alias.
	- Importação somente para efeito colateral (import _).
	- Organização de imports em grupos:
		1. Pacotes padrão
		2. Pacotes de terceiros
		3. Pacotes internos do projeto
*/

import (
	"fmt"              // pacote padrão para saída formatada
	"math"             // pacote padrão com funções matemáticas
	_ "net/http/pprof" // importação para efeito colateral (ativa pprof)
	t "time"           // importação com alias (apelido)
)

func main() {
	fmt.Println("Demonstração de imports")

	// Usando pacote math
	raiz := math.Sqrt(64)
	fmt.Println("Raiz quadrada de 64 é:", raiz)

	// Usando pacote time com alias
	agora := t.Now()
	fmt.Println("Hora atual:", agora.Format("15:04:05"))
}
