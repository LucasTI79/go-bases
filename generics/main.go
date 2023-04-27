package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// keyword functionname extraParams arguments returntype

// constraints
type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T Number](number1, number2 T) T {
	return number1 + number2
}

func SomaAny[T comparable](number1, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

func Max[T constraints.Ordered](number1, number2 T) T {
	if number1 > number2 {
		return number1
	}

	return number2
}

type stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	fmt.Println(concat([]MyString{"a", "b", "c"}))

	var x, y, z MyNumber
	x = 1
	y = 2
	z = 2

	fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
	fmt.Println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SomaGenerica(map[string]float64{"a": 1, "b": 2, "c": 3}))
}
