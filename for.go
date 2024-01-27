package main

import (
	"fmt"
	"time"
)

const nVezes = 6

func main(){
	respArrNames := arrayStates()

	// for classico

	// for i := 0; i < len(respArrNames); i++ {
	// 	name := respArrNames[i]
	// 	fmt.Println(i,"- nome: "+name)
	// }

	// for Enxulto usando o range

	// for i, name := range respArrNames {
	// 	fmt.Println("Posição:",i,"nome:",name)
	// }

	
	// Contanto a quantidade de for

	for i :=0; i < nVezes; i++ {
		fmt.Println(i,"==========================")

		for i, name := range respArrNames {
			fmt.Println("Posição:",i,"nome:",name)
		}

		time.Sleep(5 * time.Second)
	}

}


func arrayStates()[]string{
	names := []string{}
	
	names = append(names, "João", "Paulo", "Daniela", "Patrícia")
	
	return names
}
