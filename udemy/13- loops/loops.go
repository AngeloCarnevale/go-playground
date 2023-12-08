package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
	}

	nomes := [3]string{"Ângelo", "Diego", "Nicole"}

	for index, nome := range nomes {
		fmt.Println (index, nome)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}