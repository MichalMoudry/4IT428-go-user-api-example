package model

import "time"

type User struct {
	Email     string
	FirstName string
	LastName  string
	BirthDate DateTime
}

type DateTime struct {
	Day   int
	Month time.Month
	Year  int
}
