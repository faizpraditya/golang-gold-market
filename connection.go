package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// func goDotEnvVariable(key string) string {

// 	// load .env file
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

// func envVar() string {

// 	// load .env file
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	datasourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"))

// 	return datasourceName
// }

// func connectDB(datasourceName string) {

// 	// Connect to databasae
// 	sqlx.Connect("pgx", datasourceName)
// 	db, err := sqlx.Connect("pgx", datasourceName)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println("Successfully connect to database!")
// 	}

// 	// Close connection to database
// 	defer func(db *sqlx.DB) {
// 		err := db.Close()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}(db)

// }

type Database struct {
	conn *sqlx.DB
	err  error
}

func connectDB() *Database {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	datasourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Connect to databasae
	db, err := sqlx.Connect("pgx", datasourceName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connect to database!")
	}
	return &Database{db, err}
}

// func closeDB() {
// 	// Close connection to database
// 	defer func(db *sqlx.DB) {
// 		err := db.Close()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}(db)
// }
