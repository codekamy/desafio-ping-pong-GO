//  CANAL : modo de duas goroutines se comunicarem e sincronizarem sua execução

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

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)    // delay para o print
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
