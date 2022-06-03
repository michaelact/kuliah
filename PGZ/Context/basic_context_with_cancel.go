package main

import (
	"runtime"
	"context"
	"time"
	"fmt"
)

func GenerateNumber(ctx context.Context) chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)

		counter := 1
		for {
			select {
			case <-ctx.Done():
				// Prevent go routine leaks
				return
			default:
				channel <- counter
				counter++
			}
		}
	}()

	return channel
}

func main() {
	fmt.Println("Total Go Routines (Start): ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	channel := GenerateNumber(ctx)
	fmt.Println("Total Go Routines (In Progress): ", runtime.NumGoroutine())

	for number := range channel {
		fmt.Println(number)
		if number == 10 {
			break
		}
	}

	// Put to sleep after sending a cancel signal to wait for the cancellation process
	cancel()
	time.Sleep(time.Second * 2)

	fmt.Println("Total Go Routines (Done): ", runtime.NumGoroutine())
}
