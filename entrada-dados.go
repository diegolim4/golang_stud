package main

import "fmt"

func main(){
	nome := "Diego Lima"
	versao := 1.1

	fmt.Println("Olá ", nome)
    fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	/* 
		Para captura a entrada usamos a função do pacote fmt Scanf()
		Ele recebe dois argumentos um modificador (string, inteiro etc..) e um 
		ponteiro que seria a variavel  
	*/

	var comando int
	// O & é para pegar o endereço/ponteiro de onde esta alocado o valor da variavel
	// fmt.Scanf("$d", &comando)

	fmt.Scan(&comando) // Só o Scan() não precisamos passar o modificador

	fmt.Println("O endereço da variavel comando: ", &comando) // Vai retorna um valor hexadecimal

}