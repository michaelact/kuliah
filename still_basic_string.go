package main

import (
    "golang.org/x/exp/slices"
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

    vowels := []string{ "a", "i", "u", "e", "o" }
    for _, chr := range input.Text() {
        if !slices.Contains(vowels, string(chr)) {
            fmt.Printf("%c", chr)
        }
    }
}
