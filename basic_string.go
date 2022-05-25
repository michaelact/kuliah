package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    input := bufio.NewScanner(os.Stdin)

    // Get Input
    fmt.Printf("Masukkan kalimat = ")
    input.Scan()
    if input.Err() != nil {
        fmt.Println(input.Err())
    }

    fmt.Println("Panjang kalimat =", len(input.Text()))
}
