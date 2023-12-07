package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	
	p1 := pessoa{"Ângelo", 17, 165}
	fmt.Println(p1)

	e1 := estudante{p1, "Desenvolvimento de Sistemas", "Senai Roberto Mange"}

	fmt.Println(e1)
}
