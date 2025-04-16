package main

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

	DBCreate := `
	CREATE TABLE public.test (
	id integer,
	name character varying COLLATE pg_catalog."defaults"
	)
	WITH (
	OIDS = FALSE
	)
	`

	_, err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("table created")
	}
	db.Close()
}
