package main

import (
	"fmt"
	"reflect"
)

func main(){

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Pizza")
	fmt.Println("2 - Hamburguer")

	var opt int

	fmt.Scan(&opt) 
	fmt.Println("type ===>", reflect.TypeOf(opt).Kind() != reflect.Int)

	 /* 
	 Para fazer a validação do tipo usamos o Kind() para fazer a verificação subjacente
	 Oq seria subjacente ?
	 Em Go, int, int8, int16, int32, int64 são tipos diferentes, a verificão acima é 
	 somente para int (subjacente)  
	 */

	// No go o if sempre tem que ser uma instruçao que retorne um boolean
	if opt == 1 {
		fmt.Println("Vc escolheu a opção", opt, "que é pizza")
	}else if opt == 2 {
		fmt.Println("Vc escolheu a opção", opt, "que é Hamburguer")

	}else{
		fmt.Println("Vc escolheu a opção", opt, "que é INVÁLIDA!")

	}


}