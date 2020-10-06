package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Customer struct {
	ID        int    `json:"customer_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our db connection.
	// The db is called testdb.
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close until after main() has finished
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT customer_id, first_name, last_name FROM customers")

	for results.Next() {
		var customer Customer

		// for each row, scan the result into our customer composite object
		err = results.Scan(&customer.ID, &customer.FirstName, &customer.LastName)
		if err != nil {
			panic(err.Error())
		}

		log.Printf(customer.FirstName)
		log.Printf(customer.LastName)
	}
}
