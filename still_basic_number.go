package main

import (
    "fmt"
)

func main() {
    result := 0;
    // Initialize number of array
    numbers := []int { 3, 4, 10, 11, -1, 10, 22 }
    for _, val := range numbers {
        result += val
    }

    fmt.Println(result)
}
