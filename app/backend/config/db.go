package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the database handler for handle all db operations
var DB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "../../files/doneit.db")
	if err != nil {
		panic(err)
	}

	DB = db
}
