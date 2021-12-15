package main

import (
	"time"
)

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
	// DomisiliID  int       `db:"domisili_id"`
	DomisiliID Domicile
	IsActived  int `db:"is_actived"`
}

type Domicile struct {
	Id           int
	DomicileName string `db:"domisili"`
}

type CustomerDomicile struct {
	Id           int
	FirstName    string  `db:"first_name"`
	LastName     string  `db:"last_name"`
	DomicileName string  `db:"domisili"`
	Count        int     `db:"count"`
	Avg          float32 `db:"avg"`
}
