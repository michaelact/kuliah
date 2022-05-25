package main

import (
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
	"io"
)

func BubbleSort(arr []int64) []int64 {
	length := len(arr)

	// Initial Value is 1 to match the for loop condition
	swapCount := 1

	for swapCount > 0 {
		swapCount = 0
		for i := 0; i < len(arr); i++ {
			j := i+1
			// Ensure 'j' value is under array length
			if j < length {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
					swapCount++
				}
			}
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
	arr = BubbleSort(arr)
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
