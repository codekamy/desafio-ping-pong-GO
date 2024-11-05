// Goroutines : é uma função capaz de executar de modo concorrente com outras funções

package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)   // go é a palavra reservada 
	var escreva string
	fmt.Scanln(&escreva)        // cria uma nova linha, varre o texto lido da entrada padrão, armazenando valores sucessivos
}
