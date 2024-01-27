package main 

import (
	"fmt"
)

func main() {

	// testArray()
	
	// arrStates := returnStates()
	// fmt.Println(arrStates)

	testSlice()
}

func testArray(){
	// Um array no Go é limitado ao quantidade que colocamos dentro do array
	var arr [4] string

	arr[0] = "Maça"
	arr[1] = "Uva"
	arr[2] = "Banana"
	arr[3] = "Laranja"

	fmt.Println(arr)

	// Por essa limitação, em Go geralmente não trabalhamos com arrays, 
	// e sim com uma outra estrutura de dados, chamada Slice, que
	// funciona em cima do array, mas não tem tamanho fixo.u
}

func returnStates() [5]string{
	var estados [5] string
	estados[0] = "SP"
	estados[1] = "RJ"
	estados[2] = "MG"
	estados[3] = "ES"
	estados[4] = "RS"

	return estados
}

func testSlice(){
// Slices, são abstrações do array em Go, o usamos em cima do arrays

// Declarando o slice
nomes := []string{"John", "Paul", "Geoger"}
fmt.Println(nomes)
fmt.Println("Tamanho nomes: ", len(nomes))
fmt.Println("Capacidade nomes: ", cap(nomes))

// Colocando um valor dentro de slice

nomes = append(nomes, "Ringo")
fmt.Println("Dps do append:",nomes)
fmt.Println("Tamanho nomes, dps do append:", len(nomes))
fmt.Println("Capacidade nomes, dps do append:", cap(nomes)) // Após o append a capacidade do array dobra

}