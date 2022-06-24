package main

import (
	"fmt"
)

type Car[T any] struct {
	First  T
	Second T
}

// Gunakan underscore jika tidak ingin menggunakan type generic
func (self *Car[_]) GetFirstCar(message string) {
	fmt.Println(message, self.First)
}

func (self *Car[T]) ChangeSecondCar(newCar T) T {
	self.Second = newCar
	return self.Second
}

func main() {
	// Type inference tidak bisa digunakan untuk struct
	myCar := Car[string]{
		First: "Ferrari", 
		Second: "Avanza", 
	}

	myCar.GetFirstCar("Hi! It's my car")
	myCar.ChangeSecondCar("Honda")
	fmt.Println(myCar.Second)
}
