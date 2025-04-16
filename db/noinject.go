package main
// skip the whole creation process

insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")

if err != nil {
	panic(err)
}

_, err = insert.Exec(2, "second")
if err != nil {
	panic(err)
}
fmt.Println("value inserted")
defer db.Close()
}
