package main

import "fmt"

type user struct {
	name     string
	email    string
	password string
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     string
}

func main() {
	var u user
	u.name = "Tião"
	u.email = "tiao@email.com"
	u.password = "1234"

	fmt.Println(u)

	usuario2 := user{"Ângelo", "angelo@email.com", "123", endereco{logradouro: "Rua Para Lá de Bom Grado", numero: "420"}}
	fmt.Println(usuario2)

	usuario3 := user{name: "Davi"}
	fmt.Println(usuario3)

}
