package main

import (
	"fmt"
	"os"
	"io/ioutil" // Com o pacote io/ioutil, chamamos a função ReadFile, para ler um arquivo
	"bufio" // Essa pacote nos ajuda a ler os conteudo dos arquivos linha a linha
	"io"
	"strings"
)	

func main(){
	// readFileOs()

	// readFileIoutil()

	readFileBufio()
}

func readFileOs(){
	// Aqui vai retorna um ponteiro que onde esta o arquivo em hexadecimal
	resFile, err := os.Open("file-test.txt") 
	
	if err != nil {
		fmt.Println("Houve um error: ",err)
		return
	}

	fmt.Println(resFile)
}

func readFileIoutil(){
	
	// Com o ioutil ele retorna um arr de bytes, por isso temos que converter para string

	byteFile, err := ioutil.ReadFile("file-test.txt") 

	if err != nil {
		fmt.Println("Houve um error: ",err)
		return
	}

	stringFile := string(byteFile) // Aqui convertemos o array byte com a função string()

	fmt.Println(stringFile)
}

func readFileBufio(){
	resFile, err := os.Open("file-test.txt")

	if err != nil {
		fmt.Println("Houve um erro - os: ", err)
	}

	// Do bufio usamos a função NewReader() que vai receber o arquivo lido
	leitor := bufio.NewReader(resFile)

	/* 

	Com ReadString() precisamos definir um delimitador para saber atpe onde deve ser lido a linha
	No nosso caso, queremos ler a linha inteiro, ou seja, até a quebra de linha, que é 
	representada pelo \n. Como queremos representar um byte, utilizaremos aspas simples
	
	*/
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) // LImpa o espaco da quebra de linha
		fmt.Println(linha)

		if err == io.EOF { // EOF (End Of File) Quando não tem mais linha para ler no arquivo
			break
		}

		if err != nil {
			fmt.Println("Houve um erro - os: ", err)
		}
		
	}
	
	// Se abrimos um arquivo com os.Open, após lê-lo, é uma boa prática fechá-lo com a função Close()
	resFile.Close()
}