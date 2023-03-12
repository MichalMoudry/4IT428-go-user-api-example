package dto

type User struct {
	Email     string    `json:"email" validate:"email"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name"`
	BirthDate BirthDate `json:"birth_date"`
}

type BirthDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
