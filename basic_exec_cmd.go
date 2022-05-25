package main

import (
    "os"
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("ip", "add")
    stdout, err := cmd.Output()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println(string(stdout))
}
