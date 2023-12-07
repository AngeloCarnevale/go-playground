package main

import (
	"fmt"
)

func main() {

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	slice := []int{}
	fmt.Println(slice)
	
	slice = append(slice, 1)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays internos
	slice3 := make([]float32, 10, 15) // Criando slice com 10 items de capacidade 15
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Capacidade do slice
}