package main

import (
	"strings"
	"strconv"
	"bufio"
    "fmt"
    "os"

    "github.com/gocarina/gocsv"
    "golang.org/x/exp/slices"
)

var (
	DRINK_NAME_LIST  = []string{ "kopi", "coklat", "soda" }
	DRINK_SIZE_LIST  = []string{ "small", "medium", "largest" }
	DRINK_SERVE_LIST = []string{ "panas", "hangat", "dingin" }
)

type Drink struct {
	ID    int    `csv:"id"`
	Name  string `csv:"nama_kopi"`
	Size  string `csv:"ukuran"`
	Serve string `csv:"sajian"`
	Price int    `csv:"harga"`
}

func readCsvFile(data interface{}, filePath string) {
	in, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer in.Close()
	if err := gocsv.UnmarshalFile(in, data); err != nil {
		panic(err)
	}
}

func main() {
	// Read previous data
	transactions := []Drink{}
	readCsvFile(&transactions, "./transaction_database.csv")

	fmt.Println("Program kasir kafe!!!")
	fmt.Println("Silahkan pilih salah satu nomor dibawah ini:")
	fmt.Println("1) Input Data")
	fmt.Println("2) View History")
	fmt.Println("3) Delete History")
	fmt.Println("4) Exit")

	input := bufio.NewScanner(os.Stdin)
	option := 0
	for option != 4 {
		fmt.Printf("Pilih menu: ")
		input.Scan()
		checkError(input.Err())

		option, err := strconv.Atoi(input.Text())
		checkError(err)

		length := len(transactions)
		lastId := transactions[length-1].ID
		if option == 1 {
			fmt.Printf("Nama minuman: ")
			input.Scan()
			drink_name := strings.ToLower(input.Text())

			fmt.Printf("Ukuran minuman: ")
			input.Scan()
			drink_size := strings.ToLower(input.Text())

			fmt.Printf("Penyajian: ")
			input.Scan()
			drink_serve := strings.ToLower(input.Text())

			// Ensure attribute names is one of the list
			if (slices.Contains(DRINK_NAME_LIST, drink_name)) && (slices.Contains(DRINK_SIZE_LIST, drink_size)) && (slices.Contains(DRINK_SERVE_LIST, drink_serve)) {
				fmt.Printf("Ketik 'yes' bila sudah yakin ingin memasukkan transaksi ini: ")
				input.Scan()
				confirmation := input.Text()
				if confirmation == "yes" {
					transaction := Drink{
						ID: lastId + 1, 
						Name: drink_name,
						Size: drink_size,
						Serve: drink_serve,
						Price: len(drink_size) * len(drink_name) * len(drink_serve) * 100 ,
					}

					transactions = append(transactions, transaction)
				}
			} else {
				fmt.Println("Ulangi kembali.")
			}
		} else if option == 2 {
			fmt.Println(transactions)
		} else if option == 3 {
			fmt.Printf("Masukkan ID transaksi yang ingin dihapus: ")
			input.Scan()
			rmId, err := strconv.Atoi(input.Text())
			checkError(err)

			// Remove specific element of transactions array
			found := false
			for i, transaction := range transactions {
				if transaction.ID == rmId {
					found = true
					transactions = append(transactions[:i], transactions[i+1:]...)
					break
				}
			}

			if !found {
				fmt.Println("ID tidak ditemukan.")
			}
		} else if option == 4 {
			break
		} else {
			fmt.Println("Pilihan tidak ditemukan.")
			continue
		}

		// Reset 'option' variable value
		option = 0
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
