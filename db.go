package feed2sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables() {
	db, err := sql.Open("sqlite3", "feed.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "ARTICLES"`)
	if err != nil {
		panic(err)
	}
}
