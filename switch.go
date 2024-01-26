package main

import (
	"fmt"
	"reflect"
)

func main() { 
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var opt int

	fmt.Scan(&opt)

	switch opt {
	case 1:
		fmt.Println("Monitorando...")
		// Não é necessário usar o break mas se colocar o compilado não reclama

	case 2:
		fmt.Println("Exibindo logs...")

	case 0:
		fmt.Println("Saindo...")

	default:
		fmt.Println("Opçao inválida!")
	}
	

}	