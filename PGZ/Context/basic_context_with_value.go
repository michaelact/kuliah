package main

import (
	"context"
	"fmt"
)

/*
                  ___ Civic
      __ Honda __/
     /           \___ Jazz
mobil
     \__ Toyota _____ Fortuner
*/

func main() {
	mobil := context.Background()
	honda := context.WithValue(mobil, "honda", "Honda")
	toyota := context.WithValue(mobil, "toyota", "Toyota")
	civic := context.WithValue(honda, "civic", "Civic")
	jazz := context.WithValue(honda, "jazz", "Jazz")
	fortuner := context.WithValue(toyota, "fortuner", "Fortuner")

	fmt.Println(mobil)
	fmt.Println(honda)
	fmt.Println(toyota)
	fmt.Println(civic)
	fmt.Println(jazz)
	fmt.Println(fortuner)

	// .Value() function will search from parent to root
	fmt.Println(jazz.Value("honda"))
}
