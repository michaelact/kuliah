package main

import (
	"fmt"
)

const (
	Monday int = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	today := Tuesday
	fmt.Println("The day number stored in 'today' is", today)
}
