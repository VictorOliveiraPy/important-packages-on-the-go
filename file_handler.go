package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Hello, world!")
	tamanho, err := f.Write([]byte("Hello, world!")) // se nao souber que e string utilizar byte

	if err!= nil {
        panic(err)
    }
	fmt.Println("Arquivo: ", tamanho)
	f.Close()


	// leitura

	arquivo, err := os.ReadFile("arquivo.txt")
	if err!= nil {
			panic(err)
		}

		fmt.Println(string(arquivo))

		// leitura de pouco em pouco abrindo o arquivo
		readFile()
}


func readFile(){
	arquivo, err := os.Open("arquivo.txt") 
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
        if err!= nil {
            break

        }
        fmt.Println(string(buffer[:n]))
	}
}