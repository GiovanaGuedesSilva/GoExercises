
# Go Projects — Section 1: Fundamentos da Linguagem

Este diretório contém os fundamentos básicos da linguagem Go, organizados por tópicos essenciais para quem está começando. Cada subpasta traz exemplos práticos e arquivos com código comentado.

## 📂 Estrutura dos Diretórios

- `1_basics/` — Fundamentos da linguagem (variáveis, funções, condicionais, etc.)
- `2_collections/` — Trabalhando com arrays, slices, maps e strings.
- `3_structs_objects/` — Definição de structs, métodos, interfaces e princípios de design.
- `4_modules_architecture/` — Organização de código em módulos e arquitetura de software em Go.

## ▶️ Como executar os exemplos

1. Certifique-se de ter o Go instalado (https://golang.org/doc/install)
2. Acesse o diretório do exemplo desejado:
```sh
cd section_1/1_basics/01_hello-world
go run main.go
```
3. Para criar um novo módulo em uma pasta, use:
```sh
go mod init nomedomodulo
```

## ✅ Boas práticas usadas

- Cada exemplo possui um `main.go` separado e bem comentado.
- O código é simples, direto e didático.
- Nome das pastas descreve o conteúdo de forma clara.
- Uso de `go mod` para controle de dependências onde necessário.

## 🧠 Sugestões para estudo

- Leia os comentários nos arquivos `.go`.
- Execute os códigos e tente modificar os valores.
- Use o comando `go fmt` para manter a formatação do código.

---
