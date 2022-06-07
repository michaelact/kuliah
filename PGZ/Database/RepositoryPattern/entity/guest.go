package entity

import (
	"time"
)

type Guest struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
	RegDate   time.Time
}
