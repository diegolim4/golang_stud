package main

import (
	"fmt"
)

func main(){
	nameUser := getName()

	resp := greet(nameUser)

	fmt.Println("resp ", resp)

}


func getName() string {
	fmt.Println("Seja-bem vindo.")
	fmt.Println("Por favor me informe seu nome:")

	var nome string

	fmt.Scan(&nome)
	
	return nome

}

func greet(nameUser string) string {
	text := "Olá " + nameUser + ", Que bom ter você conosco!"
	return text
}
