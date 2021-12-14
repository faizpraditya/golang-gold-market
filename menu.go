package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetCustomers(db *sqlx.DB) {
	customers := []Customers{}
	db.Select(&customers, `SELECT id, first_name, last_name, date_of_birth, email FROM mst_customer order by id`)
	// log.Println(customers)

	kirana, johan, budi := customers[0], customers[1], customers[2]
	log.Printf("%#v\n%#v\n%#v\n", kirana, johan, budi)
	log.Println(customers)
}
