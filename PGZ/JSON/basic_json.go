package main

// Saat decoding JSON, sifat JSON menjadi case insensitive.
import (
    "encoding/json"
    "fmt"
    "os"
)


// Secara default, tanpa tag, JSON key akan diambil berdasarkan key struct itu sendiri.
// Misal `Name string`, maka JSON key-nya adalah Name juga.
type Student struct {
    Name            string   `json:"name"`
    Grade           string   `json:"grade"`
    Age             int      `json:"age"`
    Extracurricular []string `json:"extracurricular"`
}

func PrintJSON(data interface{}) {
    bytes, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(bytes))
}

func main() {
    // Convert Golang Data Type to JSON
    PrintJSON("Hanya")
    PrintJSON(1)
    PrintJSON(true)
    PrintJSON([]string{ "satu", "kedua", "tiga" })

    studentA := new(Student)
    studentA.Name = "Michael Act"
    studentA.Grade = "SMK"
    studentA.Age = 18
    PrintJSON(studentA)

    // Convert JSON to Golang Data Type
    studentBjson := `{"Name": "Michael Act", "Grade": "SMK", "Age": 12}`
    studentB := new(Student)
    err := json.Unmarshal([]byte(studentBjson), &studentB)
    if err != nil {
        panic(err)
    }
    fmt.Println(studentB)

    // Convert JSON to Golang Map
    var studentC map[string]interface{}
    err = json.Unmarshal([]byte(studentBjson), &studentC)
    if err != nil {
        panic(err)
    }
    fmt.Println(studentC["Name"])
    fmt.Println(studentC["Grade"])
    fmt.Println(studentC["Age"])

    // JSON Streaming Decoder
    srcfile, _ := os.Open("example.json")
    decoder := json.NewDecoder(srcfile)
    studentD := new(Student)
    decoder.Decode(&studentD)
    fmt.Println(studentD)

    // JSON Streaming Encoder
    dstfile, _ := os.Create("result.json")
    encoder := json.NewEncoder(dstfile)
    encoder.Encode(studentD)
}
