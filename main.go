package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=testdb sslmode=disable")
	rows, err := db.Query("select first_name, last_name, phone_number from users")
	columns, err := rows.Columns()

	columnHeader := strings.Join(columns, ", ")

	println(columnHeader)
	println(strings.Repeat("-", len(columnHeader)))

	for rows.Next() {
		var firstName string
		var lastName string
		var phoneNumber string
		err := rows.Scan(&firstName, &lastName, &phoneNumber)

		fmt.Printf("%s %s: %s\n", firstName, lastName, phoneNumber)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
