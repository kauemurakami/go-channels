package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)           // create chan (channel)
	go writePhrase("Hello World", channel) // go routine
	fmt.Println("Depois da função escrever ser executada")

	for { //loop infinito
		message, open := <-channel // criando message with chan value and verification is open
		if !open {                 // caso o nosso chan não esteja aberto
			fmt.Println("chan fechado estado é", open)

			break // finaliza a aplicação
		}
		fmt.Println(message, open) // output Hello world
	}

}

func writePhrase(text string, c chan string) {

	for k := 0; k < 5; k++ {
		c <- text               // atribuindo valor ao nosso chan
		time.Sleep(time.Second) // delay de um segundo
	}
	// closing the chan
	close(c) // Fecha o canal
}
