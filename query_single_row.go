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
	db, err := sql.Open("mysql", "testuser:password@tcp(127.0.0.1:3306)/testdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close until after main() has finished
	defer db.Close()

	var customer Customer

	// Execute the query
	err = db.QueryRow("SELECT customer_id, first_name, last_name FROM customers where customer_id = ?", 1).Scan(&customer.ID, &customer.FirstName, &customer.LastName)

	if err != nil {
		panic(err.Error())
	}

	log.Println(customer.FirstName)
	log.Println(customer.LastName)
}
