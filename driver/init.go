package driver

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Cats *CatsDriver
var Items *ItemsDriver
var Sessions *SessionsDriver

func init() {
	db, err := sql.Open("sqlite3", "./tutorial.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	Cats = newCatsDriver(db)
	Items = newItemsDriver(db)
	Cats.populate()
	Sessions = newSessionsDriver(db)
}
