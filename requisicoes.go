package main

import (
	"fmt"
	"net/http" //  Pacote responsável pelas requisições web
	"os"
)
// IMPORTANTE: Net é o pacote de internet do Go, o http faz parte desse pacote

func main(){
	/* 
	Por meio do httpbin.org, é possível simular o status HTTP ao passar o código de status correspondente:
	
	https://httpbin.org/status/200
	https://httpbin.org/status/404
	*/

	link := "https://google.com"

	for { // Quando usamos o for sem nenhuma instrução ele fica em loop semelhante ao while
		startingMonitoring()

		checkSite(link)
	}

	

}

func startMonitoring(){
	fmt.Println("==============================")
	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("0 - Sair")
	fmt.Println("==============================")

	
	var opt int

	fmt.Scan(&opt)

	if opt == 0 {
		fmt.Println("Saindo...")
		os.Exit(0)
	}

}

func checkSite(site string) {

	fmt.Println("Iniciando monitoramento em "+ site)
	
	response, _ := http.Get(site) // Usando _ para ignora o segundo retorno de erro


	statusCode := response.StatusCode

	fmt.Println("statusCode:",statusCode )

	if statusCode == 200 {
		fmt.Println("O Site:", site, "foi carregado com successo!")
	}else {
		fmt.Println("O Site:", site, "Esta com probemas. Status Code:", statusCode )
	}
}