package main

import (
	"sync"
	"time"
	"fmt"
)

func Counter(length int, order *Order){
	for j := 0; j < length; j++ {
		order.AddOrder(1)
	}
}

type Order struct {
	Locker sync.RWMutex
	Count  int
}

func (order *Order) AddOrder(count int) {
	order.Locker.Lock()
	order.Count += count
	order.Locker.Unlock()
}

func (order *Order) GetOrder() int {
	order.Locker.RLock()
	count := order.Count
	order.Locker.RUnlock()
	return count
}

func main() {
	order := Order{}

	for i := 0; i < 10000; i++ {
		go Counter(100, &order)
		go fmt.Println(order.GetOrder())
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasil akhir:", order.Count)
}
