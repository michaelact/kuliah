package main

import (
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
	"io"
)

func SelectionSort(arr []int64) []int64 {
	for i := 0; i < len(arr); i++ {
		smallestId := i
		// Find smallest value of array
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[smallestId] {
				smallestId = j
			}
		}

		// Swap current value with smallest value
		if (smallestId > i) {
			arr[i], arr[smallestId] = arr[smallestId], arr[i]
		}
	}

	return arr
}

func main() {
	// Get input
	fmt.Printf("Masukkan angka (dipisahkan dengan spasi) = ")
	input := bufio.NewReader(os.Stdin)
	arrTemp := strings.Split(strings.TrimSpace(readLine(input)), " ")

	// Convert string to integer
	var arr []int64
	for i := 0; i < len(arrTemp); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	fmt.Println("Before sorted: ", arr)
	arr = SelectionSort(arr)
	fmt.Println("After sorted: ", arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
