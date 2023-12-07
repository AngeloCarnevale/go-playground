package main

import (
	"errors"
	"fmt"
)

// tipos de inteiros: int8, int16, int32, int64

func main() {
	var number int = -100000000
	fmt.Println(number)

	var numero2 uint32 = 10000

	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 1299999999993.45
	fmt.Println(numeroReal1)

	numeroReal3 :=  12345.67
	fmt.Println(numeroReal3)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var booleano1 bool = false

	fmt.Println(booleano1)


	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
	
}