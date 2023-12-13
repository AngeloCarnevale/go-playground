package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal // Esperando que o canal receba um valor

	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		canal <- texto // Mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
