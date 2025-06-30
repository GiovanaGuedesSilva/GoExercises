package main

/*
	Demonstração do uso de packages (pacotes) em Go.

	- Em Go, todo arquivo pertence a um *package*.
	- O package `main` é usado para executáveis (com função `main()`).
	- Outros packages servem para organizar funcionalidades reutilizáveis.
	- Os nomes dos diretórios e dos packages geralmente coincidem.
	- Para importar um pacote local, usamos o caminho relativo ao módulo.

	Estrutura comum:
		project/
			main.go             → package main
		/mathutil/
				operacoes.go    → package mathutil

	- Para usar funções de outro package, elas devem começar com letra maiúscula (exportadas).
*/
