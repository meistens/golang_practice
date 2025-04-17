package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var prop string

	db, err := sql.Open("postgres", "user=postgres password=password host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection successful")
	}

	TableCreate := `
	CREATE TABLE Number
	(
		Number integer NOT NULL,
		Property text COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
	OIDS=FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE Number
	OWNER to Postgres
	`

	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("table created")
	}

	insert, insertErr := db.Prepare("INSERT INTO Number VALUES($1, $2")
	if insertErr != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "Even"
		} else {
			prop = "Odd"
		}

		_, err = insert.Exec(i, prop)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("number:", i, "is:", prop)
		}
	}
	insert.Close()
	fmt.Println("numbers ready")
	db.Close()
}
