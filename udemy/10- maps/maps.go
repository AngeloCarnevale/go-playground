package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Ângelo",
		"sobrenome": "Carnevale",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia de Software",
			"campus": "campus 1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome": "Áries",
	}

	fmt.Println(usuario2)
}