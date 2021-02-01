package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type change struct {
	table string
}

var db *sql.DB
var dbCache bool

func AccessDatabase() {
	db, err := sql.Open("mysql", "root:password@/labyrinth")
	if err != nil {
		log.Fatal(err)
	}
	//Set up important settings as per https://github.com/go-sql-driver/mysql/
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping() //ping the database, to make sure there's no errors
	if err != nil {
		//The database was not set up properly, so figure out what's wrong and act from there.
		//TODO: handle it properly, in other words, actually look into how to fix this
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
