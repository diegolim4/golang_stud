package main

import "fmt"

func main(){
	// Declarando uma variável em Go
	// Utilizamos a palavra var seguida no nome e do tipo
 	var nome string = "Johnny"
	var idade int = 27  // Se não declaramos nenhum valor o go vai sempre atribuir o valor 0
	var version float32 = 1.1 // Temos a opção float32 e float64 bits para caso de números maiores

	fmt.Println("Olá", nome,"! Vc tem ", idade, "Anos")
	fmt.Println("Esse programa esta na versão:",version)
}