package main

import (
	"database/sql"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

// func GetCustomers(db *sqlx.DB, dom int) {
func GetCustomers(db *sqlx.DB) {
	customers := []Customers{}
	db.Select(&customers, `SELECT id, first_name, last_name, date_of_birth, email 
	FROM mst_customer
	order by id`)

	// db.Select(&customers, `SELECT id, first_name, last_name, date_of_birth, email
	// FROM mst_customer
	// where domisili_id=$1
	// order by id`, dom)
	// log.Println(customers)

	// kirana, johan, budi := customers[0], customers[1], customers[2]
	// log.Printf("%#v\n%#v\n%#v\n", kirana, johan, budi)
	log.Println(customers)
}

func GetCustomerByID(db *sqlx.DB, id int) {
	customer := Customers{}
	err := db.Get(&customer, `SELECT Id, first_name, last_name, date_of_birth, email FROM mst_customer where id = $1 order by id`, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}

func GetCustomerByName(db *sqlx.DB, name string) {
	customer := Customers{}
	err := db.Get(&customer, `SELECT Id, first_name, last_name, date_of_birth, email FROM mst_customer where first_name=$1 order by id`, name)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}

func GetCustomerByAge(db *sqlx.DB, age int) {
	customer := Customers{}
	err := db.Get(&customer, `SELECT COUNT(id) as CountCustomerNotInJakarta
	FROM m_customer
	WHERE (extract(YEAR FROM CURRENT_DATE) - extract(YEAR FROM date_of_birth)) > 25 `, age)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}

func GetTotalCustomer(db *sqlx.DB) {
	var count int
	db.Get(&count, `SELECT COUNT(*) FROM mst_customer`)
	log.Println(count)

}

func GetCustomersDomicile(db *sqlx.DB) {
	var customers []Customers
	query := `
	SELECT mst_customer.id, mst_customer.first_name, mst_customer.last_name, mst_domisili.domisili
	FROM
	mst_customer JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	ORDER BY mst_customer.id`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		customer := Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.DomisiliID.DomicileName)

		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, customer)
	}

	for _, result := range customers {
		log.Println(result)
	}
}

func GetCustomersByEmail(db *sqlx.DB, email string) {
	customers := Customers{}
	err := db.Get(&customers, `SELECT id, first_name, last_name, date_of_birth, email FROM mst_customer WHERE email=$1`, strings.ToLower(email))

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows exist!")
		} else {
			log.Fatal(err)
		}
	}

	log.Println(customers)
}

func FindCustomerByDomicile(db *sqlx.DB, domicile string) {
	var customers []CustomerDomicile
	query := `
	SELECT mst_customer.id, mst_customer.first_name, mst_customer.last_name, mst_domisili.domisili
	FROM
	mst_customer JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	WHERE mst_domisili.domisili ilike $1
	ORDER BY mst_customer.id`

	rows, err := db.Queryx(query, domicile)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		customer := CustomerDomicile{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.DomicileName)

		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, customer)
	}

	for _, result := range customers {
		log.Println(result)
	}
}

// Prepare
// Untuk mengurangi cost reuse compare statement string, jika dipakai berkali-kali, dan banyak data
func GetCustomerByIDWithPrepare(db *sqlx.DB, id int) {
	stmt, err := db.PrepareNamed(GET_CUSTOMER_BY_ID_PREPARE)
	if err != nil {
		log.Fatal(err)
	}

	customer := CustomerDomicile{}
	customerID := map[string]interface{}{"id": id}

	err = stmt.Get(&customer, customerID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}

func UpdateCustomerByEmail(db *sqlx.DB, email string, newEmail string) {
	customer := Customers{}
	err := db.Get(&customer, `UPDATE mst_customer
	SET email=$1
	WHERE email=$2`, newEmail, email)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}

func TotalCustomerByDomicile(db *sqlx.DB) {
	// customer := Customers{}
	// err := db.Get(&customer, `SELECT mst_domisili.domisili, count(mst_customer.domisili_id)
	// FROM mst_customer
	// JOIN mst_domisili
	// ON mst_customer.domisili_id = mst_domisili.id
	// GROUP BY mst_domisili.domisili`)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(customer)

	customers := []CustomerDomicile{}
	err := db.Select(&customers, `SELECT mst_domisili.id, mst_domisili.domisili, count(mst_customer.domisili_id)
	FROM mst_customer
	JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	GROUP BY mst_domisili.id
	ORDER BY mst_domisili.id`)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customers)
}

func AgeAvgCustomerByDomicile(db *sqlx.DB) {
	customers := []CustomerDomicile{}
	// Age()
	err := db.Select(&customers, `
	SELECT mst_domisili.id, mst_domisili.domisili, avg(DATE_PART('year', current_date) - DATE_PART('year', mst_customer.date_of_birth))
	FROM mst_customer
	JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	GROUP BY mst_domisili.id
	ORDER BY mst_domisili.id`)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customers)
}

// Update lebih baik pakai NamedExec (sebagai mana fungsinya)
// Get, single
// Select, multiple
// Update, Delete NamedExec, Exec
func UpdateIsActiveById(db *sqlx.DB, id int, is_actived int) {
	customer := Customers{Id: id, IsActived: is_actived}
	_, err := db.NamedExec(UPDATE_IS_ACTIVE, customer)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Id doesn't exist")
		}
	} else {
		check_error(err, "update")
	}
}

func check_error(err error, s string) {
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Successfully " + s + " to database")
	}
}
