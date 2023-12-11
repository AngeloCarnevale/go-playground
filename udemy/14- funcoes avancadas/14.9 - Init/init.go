package main

import "fmt"

var n int

// Função init é executada antes da função main
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}