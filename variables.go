package main

import "fmt"
// O pacote reflect tem a função TypeOf, para saber o tipo
import "reflect" 

func main(){
	// Declarando uma variável em Go
	// Utilizamos a palavra var seguida no nome e do tipo
 	var nome string = "Johnny"
	var idade int = 27  // Se não declaramos nenhum valor o go vai sempre atribuir o valor 0
	var version float32 = 1.1 // Temos a opção float32 e float64 bits para caso de números maiores

	// Inferência de tipos
	var nome_2 = "Souza" // Tmbm podemos não definir o tipo que o go entende qual é o tipo
	var version_2 = 2.0

	// com o := (dois pontos e igual) o Go também entende que estamos declarando uma variável
	nome_3 := "mary" 
	
	fmt.Println("Olá", nome,"vc tem", idade, "Anos !")
	fmt.Println("Esse programa esta na versão:",version)

	// Descobrir o tipo
	fmt.Println(reflect.TypeOf(nome_2))
	fmt.Println(reflect.TypeOf(version_2)) // Quando não definimos o tipo em float o go coloca 64bit

	fmt.Println("nome_3: ", nome_3)
}