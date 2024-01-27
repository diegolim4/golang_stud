package main

import (
	"fmt"
)

func main(){
	nameUser := getName()

	resp := greet(nameUser)

	fmt.Println("resp ", resp)

	band, album, _ := favorites() // Para ignora um valor retornado usamos _
	fmt.Println(band)
	fmt.Println(album)
	// fmt.Println(year)

}

// Function para retorna dois ou mais valores
func favorites() (string,string, int) {
	nameBand := "The Beatles"
	nameAlbum := "Let it be"
	ano := 1969

	return nameBand, nameAlbum, ano
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
