//  DESAFIO:  Utilizar canais e goroutines para que o programa exiba, em alternância, as palavras ping e pong

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { // chan : palavra reservada para canal
	for i := 0; ; i++ {
		c <- "ping" // usado para enviar e receber mensagem pelo canal
	}
}
func pongar(c chan string) { // chan : palavra reservada para canal
	for i := 0; ; i++ {
		c <- "pong" // usado para enviar e receber mensagem pelo canal
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // delay para o print
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go pongar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
