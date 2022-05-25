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
    fmt.Println("PROGRAM MENENTUKAN BILANGAN GANJIL ATAU GENAP!!!")
    fmt.Printf("Masukkan angka = ")
    input.Scan()
    if input.Err() != nil {
        fmt.Println(input.Err())
    }

    // Convert input to Integer
    number, err := strconv.Atoi(input.Text())
    if number < 1 || err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    if number % 2 == 0 {
        fmt.Println(number, "adalah bilangan genap")
    } else {
        fmt.Println(number, "adalah bilangan ganjil")
    }
}
