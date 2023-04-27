package main

import "fmt"

type Carro struct {
	Name string
}

func (c Carro) andar() {
	fmt.Println(c.Name, "carro andou")
}

func main4() {
	result, str := sub(10, 3)

	resultado := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}

		return func() int {
			return res + res
		}
	}

	fmt.Println(result, str, resultado(1, 2)())

	carro := Carro{
		Name: "BMW",
	}

	carro.andar()
}

func sub(a, b int) (result int, str string) {
	return a - b, "subtraiu"
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}

	return resultado
}
