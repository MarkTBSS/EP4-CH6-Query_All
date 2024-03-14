package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	database, err := sql.Open("postgres", "user=postgres password=Pass1234 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()
	stagement, err := database.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("can't prepare query all users statment", err)
	}
	rows, err := stagement.Query()
	if err != nil {
		log.Fatal("can't query all users", err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, name, age)
	}
	fmt.Println("query all users success")
}
