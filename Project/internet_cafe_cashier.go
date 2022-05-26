// Warnet namanya kalau di Indonesia
package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
)

const (
	PRICE_PER_HOUR = 2000
)

func main() {
	fmt.Printf("Lama waktu bermain (dalam jam): ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	playTime, err := strconv.Atoi(input.Text())
	checkError(err)

	price := PRICE_PER_HOUR * playTime
	fmt.Println("Berikut harga yang harus dibayarkan:", price)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
