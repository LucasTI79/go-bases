package main

import "fmt"

func main() {
	forever := make(chan string)

	go func() {
		for {
		}
	}()

	fmt.Println("Aguardando para sempre")
	<-forever
}
