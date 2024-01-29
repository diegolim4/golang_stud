package main

import (
	"fmt"
	"os"
	"io/ioutil" // Com o pacote io/ioutil, chamamos a função ReadFile, para ler um arquivo
	"bufio" // Essa pacote nos ajuda a ler os conteudo dos arquivos linha a linha
	"io"
	"strings"
	"time"
)	

func main(){
	// readFileOs()

	// readFileIoutil()

	// readFileBufio()

	arrTexts := arrTexts()

	for _, text := range arrTexts {
		crateAndWriteInFile(text)
		
		time.Sleep(3 * time.Second)
	}

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

func crateAndWriteInFile(text string){
	/* 
		Primeiro iremos usar umas funções do os para criar um arquivo caso não exista

		OpenFile - Função do os para o Go criar o arquivo. Ela receve três parâmetros
			- Nome do arquivo
			- Flag para representar o que fazer com o arquivo
			- Tipo de permissão
			- Documentação: (doc: https://pkg.go.dev/os#pkg-constants)

		O_RDWR - Flag do OpenFile para ler e escrever no arquivo

		O_Create - Flag do OpenFile para criar um arquivo

		O_Append - Flag para que o texto seja escrito ao final do arquivo(Evita sobrescrever)

		0666 - Permisão para ler e escrever no arquivo

	*/

	file, err := os.OpenFile("text.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666) 

	if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

	file.WriteString(text + "\n")

	file.Close()
}

func arrTexts()[]string{
	texts := []string{}

	texts = append(texts, "Olá Mundo!", "Estou estudando Golang", "Tudo bem com vocês?", "Por aqui esta ok")

	return texts

}