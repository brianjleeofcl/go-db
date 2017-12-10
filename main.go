package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=testdb sslmode=disable")
	rows, err := db.Query("select * from users")
	columns, err := rows.Columns()

	columnHeader := strings.Join(columns, ", ")

	println(columnHeader)
	println(strings.Repeat("-", len(columnHeader)))

	if err != nil {
		log.Fatal(err)
	}
}
