package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const dbPort = 5433

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=postgres password=lol port=%v dbname=recipeguru sslmode=disable", dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database")
	return db, nil
}

func PrintTitles(db *sql.DB) error {
	rows, err := db.Query("SELECT title FROM ingredients")

	if err != nil {
		log.Fatal(err)
		return err
	}

	for rows.Next() {
		var title string

		err := rows.Scan(&title)
		if err != nil {
			return err
		}

		fmt.Printf("Title: %s\n", title)
	}
	return nil
}
