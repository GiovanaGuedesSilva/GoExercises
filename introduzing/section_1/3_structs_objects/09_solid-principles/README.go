package main

/*
	Demonstração dos princípios SOLID adaptados para Go.

	S - Single Responsibility Principle (Princípio da Responsabilidade Única)
		O código deve ter apenas um motivo para mudar.
		Em Go, isso significa separar responsabilidades em diferentes structs, interfaces e funções.

	O - Open/Closed Principle (Aberto/Fechado)
		Entidades devem estar abertas para extensão, mas fechadas para modificação.
		Usamos interfaces para permitir que novos comportamentos sejam adicionados sem alterar o código existente.

	L - Liskov Substitution Principle (Princípio da substituição de Liskov)
		Subtipos devem poder ser substituídos por seus tipos base sem quebrar o programa.
		Go não usa herança, mas implementações de interfaces devem respeitar esse comportamento esperado.

	I - Interface Segregation Principle (Princípio da Segregação de Interface)
		Interfaces específicas e pequenas são preferíveis a interfaces grandes e genéricas.

	D - Dependency Inversion Principle (Princípio da Inversão de Dependência)
		Código de alto nível não deve depender de implementações concretas.
		Em Go, usamos interfaces para inverter dependências.

	A seguir, exemplos práticos de cada um.
*/
