package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our db connection.
	// The db is called testdb.
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	//defer the close until after main() has finished
	defer db.Close()
}
