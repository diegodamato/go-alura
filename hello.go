package main

import (
	"fmt"
	"reflect"
	"os"
	"net/http"
)

func main() {
	

	exibeIntroducao()
	
	for{
		exibeMenu()
		
		comando := leComando()
	
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao(){
	nome := "Douglas"
	idade := 29
	versao := 1.1

	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}

func leComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	//site:="https://www.alura.com.br"
	site := "https://random-status-code.herokuapp.com"

	resp,_ := http.Get(site)
	
	if resp.StatusCode == 200{
		fmt.Println("Site", site, "foi carregado com sucesso")
	}else{
		fmt.Println("Site", site, "esta com problemas. Status Code:", resp.StatusCode)
	}

}