package main

import (
	"database/sql"
	"log"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	/* Open the database connection. */
	db, err := sql.Open("mysql", "usuario:contrasenia@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//Query insert.
	insert, err := db.Query("INSERT INTO test values (2,'TEST')")

	if err != nil {
		panic(err.Error())
	}

	//Careful making defers from queries when making transactions.
	defer insert.Close()

	//Query
	results, err := db.Query("SELECT id,name FROM tags")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag

		// For every row scan the result in a tag. The scan has to have the same amount of
		// parameters as the db if not it fails.
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		//Print Tag name.
		log.Printf(tag.Name)
	}

	var tag Tag

	//Query from a single row. The 2 is the amount of parameters it takes.
	err = db.QueryRow("SELECT id,name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error())
	}

	log.Println(tag.Name)
	log.Println(tag.ID)

}
