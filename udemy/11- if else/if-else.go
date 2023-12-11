package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15 ")
	}

	// if init: inicializa a variável para utilizar apenas dentro do escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero < -10{
		fmt.Println("Número menor que -10")
	} else {
		fmt.Println("Número está entre 0 e -10")
	}
}
