package main

import "time"

// dibuat public karena akan dipakai di package sqlx
type Customers struct {
	Id        int
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	// struct tag
	DateOfBirth time.Time `db:"date_of_birth"`
	Address     string
	Status      int
	Email       string
	Username    string    `db:"user_name"`
	Password    string    `db:"user_password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DomisiliID  int       `db:"domisili_id"`
}
