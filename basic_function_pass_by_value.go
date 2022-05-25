package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
)

func lucas(number int) int {
	if number == 0 {
		return 2
	} else if number == 1 {
		return 1
	} else {
		return lucas(number-1) + lucas(number-2)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)

    // Get Input
    fmt.Println("PROGRAM MENCARI ANGKA LUCAS BERDASARKAN POSISI!!!")
    fmt.Printf("Masukkan posisi angka yang ingin dicari = ")
    input.Scan()
    if input.Err() != nil {
        fmt.Println(input.Err())
    }

    // Convert input to Integer
    number, err := strconv.Atoi(input.Text())
    if number < 0 || err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    result := lucas(number)
   	fmt.Printf("Angka lucas pada posisi angka ke %d adalah %d", number, result)
}
