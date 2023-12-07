package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória

	var variavel3 int // Essa variável guarda um número inteiro
	var ponteiro *int // Essa variável guarda o endereço de memória de um número inteiro

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // Desreferenciação

	variavel3 = 101

	fmt.Println(variavel3, *ponteiro)

}