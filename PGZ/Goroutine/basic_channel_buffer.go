package main

import (
	"time"
	"fmt"
)

func DisplayNumber(channel chan<- string) {
	for i := 1; i < 5; i++{
		result := fmt.Sprintln("Muncul", i)
		channel <- result
	}

	fmt.Println("Selesai mengirim data.")
}

func GetNumber(channel <-chan string) {
	for i := 1; i < 5; i++ {
		go fmt.Println(<- channel)
	}

	fmt.Println("Selesai menerima data.")
}

func main() {
	channel := make(chan string, 5)
	defer close(channel)

	go DisplayNumber(channel)
	go GetNumber(channel)

	time.Sleep(5 * time.Second)
}
