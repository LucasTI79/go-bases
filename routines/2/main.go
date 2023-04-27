package main

import "fmt"

// T1
func main() {
	// T1 <> T2
	hello := make(chan string)

	// T2
	go func() {
		hello <- "hello world"
	}()

	result := <-hello

	fmt.Println(result)
}
