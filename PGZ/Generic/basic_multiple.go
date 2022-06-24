package main

import (
	"fmt"
)

func PrintNameAge[, T Number](numbers S) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(Sum([]float64{ 3.4, 3.1, 9.1 }))
	fmt.Println(Sum([]int{ 4, 4, 9 }))
}
