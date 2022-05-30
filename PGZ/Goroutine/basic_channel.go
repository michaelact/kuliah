package main

import (
	"time"
	"fmt"
)

func DisplayNumber(data int, channel chan string) {
	result := fmt.Sprintln("Muncul", data)
	channel <- result
	fmt.Println("Selesai mengirim data.")
}

func main() {
	channel := make(chan string)
	defer close(channel)
	go DisplayNumber(50, channel)

	data := <- channel
	fmt.Println("Data diterima:", data)
	time.Sleep(5 * time.Second)
}
