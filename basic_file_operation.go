package main
import (
  "golang.org/x/exp/slices"
  "io/ioutil"
  "fmt"
  "os"
)

func main() {
  data, err := ioutil.ReadFile("example_data.txt")
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  var result []byte
  vowels := []string{ "a", "i", "u", "e", "o" }
  for _, chr := range data {
    if !slices.Contains(vowels, string(chr)) {
      result = append(result, chr)
    }
  }

  err = ioutil.WriteFile("result_data.txt", result, 0600)
  if err != nil {
    os.Exit(1)
  }
}
