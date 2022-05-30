package main

import (
	"sync"
	"fmt"
)

func Counter(length int, count *int, group *sync.WaitGroup){
	for j := 0; j < length; j++ {
		defer group.Done()
		*count += 1
		group.Add(1)
	}
}

func main() {
	var group sync.WaitGroup

	count := 0
	for i := 0; i < 10000; i++ {
		go Counter(100, &count, &group)
	}

	group.Wait()
	fmt.Println(count)
}
