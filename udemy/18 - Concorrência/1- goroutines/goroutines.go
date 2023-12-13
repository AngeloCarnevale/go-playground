package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo
	// Paralelismo: tarefas executadas multithread (necessita de um processador com mais de um núcleo)
	// Concorrência: executa tarefas paralelas mesmo com um processador de um núcleo só

	go escrever("Olá mundo") // goroutine | Chama a execução mas não espera terminar para continuar o programa
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}