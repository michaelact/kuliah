package main

import (
	"strconv"
	"reflect"
	"bufio"
	"fmt"
	"os"
)

var (
	StudentAttrs = []string{ "NIM", "Name", "Subject" }

	// float64 { Weight }
	ScoreAttrs   = map[string]float64 { 
		"Present": 0.1,
		"Task":    0.2,
		"Quiz":    0.1, 
		"Forum":   0.1, 
		"Exam":    0.5,
	}

	// int{ From, To }
	GradeAttrs   = map[string][2]int{
		"A": [2]int{ 90, 100 },
		"B": [2]int{ 80, 89 },
		"C": [2]int{ 70, 79 },
		"D": [2]int{ 60, 69 },
		"E": [2]int{ 50, 59 },
	}
)

type Score struct {
	Present int
	Task    int
	Quiz    int
	Forum   int
	Exam    int
}

type Student struct {
	Id      int
	NIM     string
	Name    string
	Subject string
	Score   Score
	Grade   string
	IsPass  bool
}

func (s *Student) calculate() {
	result := 0
	for key, weight := range ScoreAttrs {
		score := getAttr(&s.Score, key).Int()
		result += int(float64(score) * weight)
	}

	for key, val := range GradeAttrs {
		if result >= val[0] && result <= val[1] {
			s.Grade = key
			break
		}
	}

	if result >= 60 {
		s.IsPass = true
	}
}

// getAttr: Copied from https://stackoverflow.com/a/66470232
func getAttr(obj interface{}, fieldName string) reflect.Value {
    pointToStruct := reflect.ValueOf(obj) // addressable
    curStruct := pointToStruct.Elem()
    if curStruct.Kind() != reflect.Struct {
        panic("not struct")
    }
    curField := curStruct.FieldByName(fieldName) // type: reflect.Value
    if !curField.IsValid() {
        panic("not found:" + fieldName)
    }
    return curField
}

func getIndexByNIM(nim string, students []Student) int {
	index := -1
	for key, student := range students {
		if student.NIM == nim {
			index = key
			break
		}
	}
	return index
}

func main() {
	// Initialize empty array of student
	var students []Student

	fmt.Println("DETERMINES STUDENT GRADUATION STATUS!!!")
	fmt.Println("Choose one option below: ")
	fmt.Println("1) Input Data")
	fmt.Println("2) Input Score")
	fmt.Println("3) View History")
	fmt.Println("4) Exit")

	input := bufio.NewScanner(os.Stdin)
	option := 0
	for option != 4 {
		fmt.Printf("Pilih menu: ")
		input.Scan()
		option, err := strconv.Atoi(input.Text())
		checkError(err)

		switch option {
		case 1:
			var student Student

			// Check student data is empty or exists
			var lastId int
			if len(students) == 0 {
				lastId = 0
			} else {
				lastId = students[len(students)-1].Id
			}

			// Get student property input
			student.Id = lastId
			for _, attr := range StudentAttrs {
				fmt.Printf("Input %s: ", attr)
				input.Scan()
				getAttr(&student, attr).SetString(input.Text())
			}

			// Check new student data is fresh or not
			if getIndexByNIM(student.NIM, students) == -1 {
				students = append(students, student)
			} else {
				fmt.Println("Data already exists, failed to input data.")
			}

		case 2:
			fmt.Printf("Input existing NIM: ")
			input.Scan()

			// Ensure NIM is already exists
			nim := input.Text()
			id := getIndexByNIM(nim, students)
			if id == -1 {
				fmt.Println("NIM not found.")
				break
			}

			// Get score input
			for attr, _ := range ScoreAttrs {
				fmt.Printf("Input %s Score: ", attr)
				input.Scan()
				tmpScore, err := strconv.ParseInt(input.Text(), 10, 64)
				checkError(err)
				getAttr(&students[id].Score, attr).SetInt(tmpScore)
			}

			// Determine graduation status and grade
			students[id].calculate()

		case 3:
			fmt.Println(students)

		case 4:
			os.Exit(0)

		default:
			fmt.Println("Selected options not found.")
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
