package main

import (
	"strings"
	"reflect"
	"strconv"
	"bufio"
	"fmt"
	"os"
)

const (
	BaseSalary = 2000000
)

var (
	EmployeeAttr = []string{ "Name", "JobClass", "Education" }
	JobWeight = map[string]float64{
		"1": 0.05, 
		"2": 0.1, 
		"3": 0.15, 
	}
	EduWeight = map[string]float64{
		"SMA": 0.025,
		"D3":  0.05,
		"S1":  0.1,
	}
)

type Salary struct {
	JobClass  int
	Education int
	Overtime  int
	Total     int
}

type Employee struct {
	Name      string
	JobClass  string
	Education string
	Salary    Salary
	WorkHours int
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

func main() {
	var employee Employee
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("EMPLOYEE SALARY CALCULATION!!!")

	// Get employee property
	for _, attr := range EmployeeAttr {
		fmt.Printf("Input %s: ", attr)
		input.Scan()
		getAttr(&employee, attr).SetString(strings.ToUpper(input.Text()))
	}

	// Get work hours
	fmt.Printf("Input Work Hours: ")
	input.Scan()
	workHours, err := strconv.Atoi(input.Text())
	checkError(err)
	employee.WorkHours = workHours

	// Check if working hours include overtime
	overHours := employee.WorkHours - 9
	if overHours > 0 {
		employee.Salary.Overtime = 3000 * overHours
	}

	employee.Salary.JobClass = int(float64(BaseSalary) * JobWeight[employee.JobClass])
	employee.Salary.Education = int(float64(BaseSalary) * EduWeight[employee.Education])
	employee.Salary.Total = BaseSalary + employee.Salary.JobClass + employee.Salary.Education + employee.Salary.Overtime

	fmt.Println(employee)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
