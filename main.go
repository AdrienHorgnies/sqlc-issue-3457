package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func main() {
	log.Println("Opening database")
	db, dbErr := sql.Open("sqlite3", "main.db")
	if dbErr != nil {
		log.Fatalln("Cannot open/create database", dbErr)
	}
	q := New(db)
	ctx := context.Background()

	log.Println("Creating tables")
	_, tablesErr := db.ExecContext(ctx, ddl)
	if tablesErr != nil {
		log.Fatalln("Cannot create tables:", tablesErr)
	}

	log.Println("Inserting users")
	insertErr := q.InsertUsers(ctx)
	if insertErr != nil {
		log.Fatalln("Cannot create users:", insertErr)
	}

	log.Println("Updating users")
	updateErr := q.UpdateUsers(ctx, "bar")
	if updateErr != nil {
		log.Fatalln("Cannot update users:", updateErr)
	}

	log.Println("Done")
}
