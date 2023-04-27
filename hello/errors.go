package main

import (
	"errors"
	"fmt"
	"log"
)

func main3() {
	res, err := soma(10, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}

func soma(x, y int) (int, error) {
	resultado := x + y
	if resultado > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return resultado, nil
}
