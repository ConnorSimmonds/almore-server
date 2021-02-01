package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type change struct {
	table string
}

var db *sql.DB
var dbCache bool

func AccessDatabase() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/labyrinth")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping() //ping the database, to make sure there's no errors
	if err != nil {
		//The database was not set up properly, so figure out what's wrong and act from there.
		//TODO: handle it properly, in other words
	}
}

func CacheChanges() {
	dbCache = true
}

func ApplyChanges() {
	if dbCache { //changes have been stored, so apply them

	} else {
		//let the client know that changes haven't even started initializing
	}
}

func CloseDatabase() {
	err := db.Close()
	if err != nil {
		//something went wrong while closing the db - look up how to do this
	}
}
