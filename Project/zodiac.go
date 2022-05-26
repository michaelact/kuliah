package main

import (
	"strings"
	"bufio"
	"time"
	"fmt"
	"os"
)

var (
	AriesFrom		= time.Date(0, 3, 21, 0, 0, 0, 0, time.Local)
	AriesTo			= time.Date(0, 4, 19, 0, 0, 0, 0, time.Local)
	TaurusFrom		= time.Date(0, 4, 20, 0, 0, 0, 0, time.Local)
	TaurusTo		= time.Date(0, 5, 20, 0, 0, 0, 0, time.Local)
	GeminiFrom		= time.Date(0, 5, 21, 0, 0, 0, 0, time.Local)
	GeminiTo		= time.Date(0, 6, 20, 0, 0, 0, 0, time.Local)
	CancerFrom		= time.Date(0, 6, 21, 0, 0, 0, 0, time.Local)
	CancerTo		= time.Date(0, 7, 22, 0, 0, 0, 0, time.Local)
	LeoFrom			= time.Date(0, 7, 23, 0, 0, 0, 0, time.Local)
	LeoTo			= time.Date(0, 8, 22, 0, 0, 0, 0, time.Local)
	VirgoFrom		= time.Date(0, 8, 23, 0, 0, 0, 0, time.Local)
	VirgoTo			= time.Date(0, 9, 22, 0, 0, 0, 0, time.Local)
	LibraFrom		= time.Date(0, 9, 23, 0, 0, 0, 0, time.Local)
	LibraTo			= time.Date(0, 10, 22, 0, 0, 0, 0, time.Local)
	ScorpioFrom		= time.Date(0, 10, 23, 0, 0, 0, 0, time.Local)
	ScorpioTo		= time.Date(0, 11, 21, 0, 0, 0, 0, time.Local)
	SagitariusFrom	= time.Date(0, 11, 22, 0, 0, 0, 0, time.Local)
	SagitariusTo	= time.Date(0, 12, 21, 0, 0, 0, 0, time.Local)
	AquariusFrom	= time.Date(0, 1, 20, 0, 0, 0, 0, time.Local)
	AquariusTo		= time.Date(0, 2, 18, 0, 0, 0, 0, time.Local)
	PiscesFrom		= time.Date(0, 2, 19, 0, 0, 0, 0, time.Local)
	PiscesTo		= time.Date(0, 3, 22, 0, 0, 0, 0, time.Local)
)

func main() {
	_currentTime := time.Now()
	time.Date(0, _currentTime.Month(), _currentTime.Day(), 0, 0, 0, 0, _currentTime.Location())
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input your name: ")
	input.Scan()
	name := input.Text()

	fmt.Printf("Input your birth date (example: 16 march): ")
	input.Scan()
	origBirth := strings.Split(input.Text(), " ")

	birthDate, birthMonth := origBirth[0], origBirth[1]
	birth, err := time.Parse("January 2", fmt.Sprintf("%s %s", birthMonth, birthDate))
	checkError(err)

	var zodiac string
	if birth.After(AriesFrom) && birth.Before(AriesTo) {
		zodiac = "Aries"
	} else if birth.After(TaurusFrom) && birth.Before(TaurusTo) {
		zodiac = "Taurus"
	} else if birth.After(GeminiFrom) && birth.Before(GeminiTo) {
		zodiac = "Gemini"
	} else if birth.After(CancerFrom) && birth.Before(CancerTo) {
		zodiac = "Cancer"
	} else if birth.After(LeoFrom) && birth.Before(LeoTo) {
		zodiac = "Leo"
	} else if birth.After(VirgoFrom) && birth.Before(VirgoTo) {
		zodiac = "Virgo"
	} else if birth.After(LibraFrom) && birth.Before(LibraTo) {
		zodiac = "Libra"
	} else if birth.After(ScorpioFrom) && birth.Before(ScorpioTo) {
		zodiac = "Scorpio"
	} else if birth.After(SagitariusFrom) && birth.Before(SagitariusTo) {
		zodiac = "Sagitarius"
	} else if birth.After(AquariusFrom) && birth.Before(AquariusTo) {
		zodiac = "Aquarius"
	} else if birth.After(PiscesFrom) && birth.Before(PiscesTo) {
		zodiac = "Pisces"
	} else {
		fmt.Println("Zodiac not found.")
		os.Exit(1)
	}

	fmt.Println("Name:", name)
	fmt.Println("Zodiac:", zodiac)
	fmt.Println("Birth Date", origBirth)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
