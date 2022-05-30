// https://en.cppreference.com/w/cpp/thread/mutex

package main

import (
	"time"
	"sync"
	"fmt"
)

func Counter(length int, count *int, mutex *sync.Mutex){
	for j := 0; j < length; j++ {
		mutex.Lock()
		*count += 1
		mutex.Unlock()
	}
}

func main() {
	var mutex sync.Mutex

	count := 0
	for i := 0; i < 10000; i++ {
		go Counter(100, &count, &mutex)
	}

	time.Sleep(5 * time.Second)
	fmt.Println(count)
}
