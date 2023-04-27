package main

import "fmt"

type Carro2 struct {
	Name string
}

func (c Carro2) andar() {
	fmt.Println(c.Name, "andou")
}

func main() {
	carro := Carro2{
		Name: "Ka",
	}

	carro.andar()
}
