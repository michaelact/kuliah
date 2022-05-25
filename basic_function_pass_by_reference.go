package main

import (
	"fmt"
)

func Change(original *string, value string) {
	fmt.Printf("Memory address of sentence after passed to the function: %p\n", original)
	*original = value
}

func main() {
	name := "michael"
	fmt.Printf("Memory address of sentence before passed to the function: %p\n", &name)
    Change(&name, "dia")
    fmt.Printf("Current Name Value = %s", name)
}
