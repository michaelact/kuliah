package main

import (
	"fmt"
)

type Fullname string
type Student struct {
	Name Fullname
	Age  uint8
}

func main() {
	student := Student {
		Name: "Michael Act", 
		Age:  18, 
	}

	fmt.Println(student)
}
