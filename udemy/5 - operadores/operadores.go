package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Operadores de Atribuição
	var variavel string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel, variavel2)

	// Operadores Relacionais (retornam booleano)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	
	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//Operadores unários
	numero := 10
	numero++
	numero += 14
	fmt.Println(numero)

	var texto string

	if numero > 5 {
		texto += "Maior que 5"
	} else {
		texto += "Menor que 5"
	}
	fmt.Println(texto)
	
}