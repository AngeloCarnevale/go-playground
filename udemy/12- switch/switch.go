package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	default:
		return "Número inválido"
	}
}


func main() {
	dia := diaDaSemana(2)

	fmt.Println(dia)
}
