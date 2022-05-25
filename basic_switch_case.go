package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input := bufio.NewScanner(os.Stdin)

    // Get Input
    fmt.Println("PROGRAM REKOMENDASI MOTOR!!!")
    fmt.Printf("Tingkat sekolah = ")
    input.Scan()
    if input.Err() != nil {
        fmt.Println(input.Err())
    }

    switch strings.ToUpper(input.Text()) {
    case "SMP":
        fmt.Println("beat street boleh sih")
    case "SMA":
        fmt.Println("satria fu cocok buat lu")
    case "KULIAH":
        fmt.Println("motor ceber mantep")
    default:
        fmt.Println("belum didukung")
    }
}
