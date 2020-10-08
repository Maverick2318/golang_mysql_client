package main

import (
        "database/sql"
        // "fmt"
        _ "github.com/go-sql-driver/mysql"
        // "io"
        "log"
        "net/http"
)

type Customer struct {
        ID        int    `json:"customer_id"`
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("This is an example server.\n"))
        // fmt.Fprintf(w, "This is an example server.\n")
        // io.WriteString(w, "This is an example server.\n")
}

func QuerySingleRow(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "text/plain")

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

        w.Write([]byte("Making your DB query.\n"))
        w.Write([]byte("First Name: " + customer.FirstName + "\n"))
        w.Write([]byte("Last Name: " + customer.LastName + "\n"))
}

func main() {
        http.HandleFunc("/hello", HelloServer)
        http.HandleFunc("/getCustomer", QuerySingleRow)
        err := http.ListenAndServeTLS(":444", "server.crt", "server.key", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
