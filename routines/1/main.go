package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(tipo, i)
	}
}

func main() {
	go contador("a")
	go contador("b")
	time.Sleep(time.Second * 10)
}
