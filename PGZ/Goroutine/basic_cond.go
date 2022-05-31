package main

import (
	"sync"
	"fmt"
)

func WaitCondition(value int, cond *sync.Cond, group *sync.WaitGroup) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Muncul", value)
	cond.L.Unlock()
	group.Done()

}

func main() {
	var group sync.WaitGroup
	var mutex sync.Mutex
	cond := sync.Cond{
		L: &mutex, 
	}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go WaitCondition(i, &cond, &group)
	}

	go func() {
		for i := 0; i < 100; i++ {
			// Key to pass the sync.Cond.Wait() function
			cond.Signal()
		}
	}()

	group.Wait()
	fmt.Println("Selesai.")
}
