package main

import (
	"time"
	"fmt"
)

func main() {
	for i := 0; i < 1000; i++ {
		// Asynchronous
		go fmt.Println(i)
	}

	time.Sleep(5 * time.Second)
}
