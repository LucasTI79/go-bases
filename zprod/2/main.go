package main

import (
	"fmt"
	"time"
)

// func contador(tipo string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(tipo, "i", i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

func main() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			queue <- i
			time.Sleep(time.Second * 1)
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}
}
