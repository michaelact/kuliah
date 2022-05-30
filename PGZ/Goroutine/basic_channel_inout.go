package main

import (
	"time"
	"fmt"
)

func OnlyIn(channel chan<- string) {
	channel <- "Nih data buatmu."
	fmt.Println("Selesai mengirim data.")
}

func OnlyOut(channel <-chan string) {
	fmt.Println("Data diterima:", <-channel)
}

func main() {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
