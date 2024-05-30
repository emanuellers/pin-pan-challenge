package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)
	go pongar(c)

	var entrada string
	fmt.Scanln(&entrada)
}

func pingar(c chan string) { // palavra reservada para Canal: chan
	for i := 0; ; i++ {
		c <- "ping" //usado para enviar e receber mensagem pelo canal

	}
}

func pongar(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}

func pin_pan() {
	for end := 0; end <= 100; end++ {
		if end%3 == 0 {
			fmt.Println(end, "Pin")
		} else if end%5 == 0 {
			fmt.Println(end, "Pan")
		}

	}
}
