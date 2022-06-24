package main

import (
	"fmt"
)

type Number interface {
	int | float32 | float64
}

func Sum[T Number](firstNum T, secondNum T) T {
	return firstNum + secondNum
}

func main() {
	fmt.Println(Sum(3, 4))
	fmt.Println(Sum(3.4, 4.5))
}
