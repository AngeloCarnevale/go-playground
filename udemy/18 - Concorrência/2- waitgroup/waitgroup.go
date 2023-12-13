package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // Quantidade de goroutines que faze parte do grupo de espera

	go func () {
		escrever("Olá mundo")
		waitGroup.Done() // -1 do waitGroup
	}() 

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 do waitGroup
	}()

	go func() {
		escrever("GoRoutine 3")
		waitGroup.Done() // -1 do waitGroup
	}()

	go func() {
		escrever("GoRoutine 4")
		waitGroup.Done() // -1 do waitGroup
	}()

	waitGroup.Wait() // 2
	
}

func escrever(texto string) {
	for i:= 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}