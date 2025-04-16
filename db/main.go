package main

// for future purposes, use docker/podman, this is here because you
// spent hours on trying to fix an - was there even one? - issue
// in the name of using podman gui over cli
// honestly, anything that makes using em easier, not make me feel
// like an insufferable tool...
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=password host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection successful")
	}

	// ping db
	connectivity := db.Ping()
	if connectivity != nil {
		panic(err)
	} else {
		fmt.Println("still connected and ready for req...")
	}

	// close connection
	db.Close()
}
