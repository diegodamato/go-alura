package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	idade := 29
	versao := 1.1

	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int

	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Não conheço esse comando")
	}

}
