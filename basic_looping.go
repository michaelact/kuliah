package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

// Factorial
func main() {
    input := bufio.NewScanner(os.Stdin)

    // Get Input
    fmt.Println("PROGRAM MENGHITUNG FAKTORIAL!!!")
    fmt.Printf("Masukkan angka = ")
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

    result := 1
    for i := 1; i <= number; i++ {
        result *= i;
    }

    fmt.Println(result)
}
