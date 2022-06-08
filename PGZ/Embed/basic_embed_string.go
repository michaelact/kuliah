package main

import (
	"fmt"
	_ "embed"
)

//go:embed testing.txt
var text string

func main() {
	fmt.Println(text)
}
