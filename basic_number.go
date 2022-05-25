package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input := bufio.NewScanner(os.Stdin)

    // Get Input
    fmt.Printf("Banyak bilangan ganjil = ")
    input.Scan()
    if input.Err() != nil {
        fmt.Println(input.Err())
    }

    // Convert input to Integer
    maxOdd, err := strconv.Atoi(input.Text())
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    for val, count := 1, 1; count <= maxOdd; val++ {
        // Check is the 'val' variable divisible by number 2
        if val % 2 != 0 {
            fmt.Printf("%d ", val)
            count++
        }
    }
}
