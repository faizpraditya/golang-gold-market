package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

// golang sebenarnya sudah provide golang database sql (interfacing)
// sqlx
// how to connect with postgres, libpq, pgx. Prefer pgx.

// go get github.com/jackc/pgx
// go get github.com/jmoiron/sqlx

// // dibuat public karena akan dipakai di package sqlx
// type Customers struct {
// 	Id        int
// 	FirstName string `db:"first_name"`
// 	LastName  string `db:"last_name"`
// 	// struct tag
// 	DateOfBirth time.Time `db:"date_of_birth"`
// 	Address     string
// 	Status      int
// 	Email       string
// 	Username    string    `db:"user_name"`
// 	Password    string    `db:"user_password"`
// 	CreatedAt   time.Time `db:"created_at"`
// 	UpdatedAt   time.Time `db:"updated_at"`
// 	DomisiliID  int       `db:"domisili_id"`
// }

func main() {
	/*
		sqlx.connect postgres
		  1. Buat konfigurasi koneksi ke database postgresql
	*/

	// Loading config database / bisa menggunakan environment (env)
	// -d -p itu namanya flag
	// .env biasanya di git ingore, hanya untuk membantu proses development
	// Api key jangan di push
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbName := "gold_market_db"
	// dbUser := "postgres"
	// dbPassword := "12345678"
	// godotenv package

	// Database connection string
	// 1. postgres://user:password@host:port/db?sslmode=disable
	// 2. host= user= dbname= sslmode=disable password= port=
	// datasourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println(stringToDate("2020-03-20"))
	// File load env harusnya diberi kondisi agar bisa dioverride oleh env var dari terminal
	db := connectDB()
	var listCustomers []Customers
	rows, _ := db.conn.Queryx("SELECT * FROM mst_customer")
	for rows.Next() {
		var c Customers
		err := rows.StructScan(&c)
		if err != nil {
			log.Fatal(err)
		} else {
			listCustomers = append(listCustomers, c)
		}
	}

	time := stringToDate("2020-11-09")
	fmt.Println(time)

	// Connect to databasae
	// sqlx.Connect("pgx", datasourceName)
	// db, err := sqlx.Connect("pgx", datasourceName)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully connect to database!")
	// }

	// // Close connection to database
	// defer func(db *sqlx.DB) {
	// 	err := db.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }(db)

	// Modifying data to database

	// Insert
	// customers := []Customers{
	// 	{
	// 		Id:          2,
	// 		FirstName:   "Faiz",
	// 		LastName:    "Praditya",
	// 		DateOfBirth: time.Date(1998, 02, 03, 10, 20, 0, 0, time.UTC),
	// 		Address:     "Semarang",
	// 		Status:      1,
	// 		Email:       "faiz@gmail.com",
	// 		Username:    "faizpraditya",
	// 		Password:    "pw",
	// 		DomisiliID:  1,
	// 		CreatedAt:   time.Now(),
	// 		UpdatedAt:   time.Now(),
	// 	}, {
	// 		Id:          3,
	// 		FirstName:   "Faiz",
	// 		LastName:    "Praditya",
	// 		DateOfBirth: time.Date(1998, 02, 03, 10, 20, 0, 0, time.UTC),
	// 		Address:     "Semarang",
	// 		Status:      1,
	// 		Email:       "faiz@gmail.com",
	// 		Username:    "faizpraditya",
	// 		Password:    "pw",
	// 		DomisiliID:  1,
	// 		CreatedAt:   time.Now(),
	// 		UpdatedAt:   time.Now(),
	// 	},
	// }
	// _, err = db.NamedExec(`INSERT INTO mst_customer
	// (id, first_name, last_name, date_of_birth, address, status, email, user_name, user_password, domisili_id, created_at, updated_at)
	// VALUES (:id, :first_name, :last_name, :date_of_birth, :address, :status, :email, :user_name, :user_password, :domisili_id, :created_at, :updated_at)`, customers)

	// newCustomer := map[string]interface{}{
	// 	"id":         4,
	// 	"first_name": "Valention",
	// 	"last_name":  "Rosa",
	// 	"status":     1,
	// 	"address":    "Grogol",
	// }

	// _, err = db.NamedExec(`INSERT INTO mst_customer
	// (id, first_name, last_name, status, address)
	// VALUES (:id, :first_name, :last_name, :status,:address)`, newCustomer)

	// _, err = db.NamedExec(`INSERT INTO mst_customer
	// (id, first_name, last_name, date_of_birth, address, status, email, user_name, user_password, domisili_id, created_at, updated_at)
	// VALUES (:id, :first_name, :last_name, :date_of_birth, :address, :status, :email, :user_name, :user_password, :domisili_id, :created_at, :updated_at)`, &Customers{
	// 	Id:          1,
	// 	FirstName:   "Faiz",
	// 	LastName:    "Praditya",
	// 	DateOfBirth: time.Date(1998, 02, 03, 10, 20, 0, 0, time.UTC),
	// 	Address:     "Semarang",
	// 	Status:      1,
	// 	Email:       "faiz@gmail.com",
	// 	Username:    "faizpraditya",
	// 	Password:    "pw",
	// 	DomisiliID:  1,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully insert data to database!")
	// }

	// Update
	// customerUpdate := Customers{
	// 	FirstName:   "Faizzz",
	// 	LastName:    "Pradityaaa",
	// 	DateOfBirth: time.Date(1998, 02, 03, 10, 20, 0, 0, time.UTC),
	// 	Address:     "Semarang",
	// 	Status:      1,
	// 	Email:       "faiz@gmail.com",
	// 	Username:    "faizpraditya",
	// 	Password:    "pw",
	// 	DomisiliID:  1,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// 	Id:          4,
	// }

	// _, err = db.NamedExec(`UPDATE mst_customer SET first_name=:first_name, last_name=:last_name, date_of_birth=:date_of_birth, address=:address, status=:status, email=:email, user_name=:user_name, user_password=:user_password, domisili_id=:domisili_id, created_at=:created_at, updated_at=:updated_at`, customerUpdate)

	// Delete
	// _, err = db.NamedExec(`DELETE FROM mst_customer WHERE id=:id`, map[string]interface{}{"id": 4})
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully delete data to database!")
	// }

	// Delete where
	// db.MustExec("DELETE FROM mst_customer WHERE id=$1", 3)
	// Delete all
	// db.MustExec("DELETE FROM mst_customer")
}
