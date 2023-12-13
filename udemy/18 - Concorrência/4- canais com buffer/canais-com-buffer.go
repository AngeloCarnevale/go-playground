package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Especifica tamanho de 2 para o canal

	canal <- "Olá mundo"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}