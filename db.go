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

	cmd := `CREATE TABLE IF NOT EXISTS "articles"(
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "url" TEXT NOT NULL UNIQUE,
		"title" TEXT,
		"read_at" TIMESTAMP,
        "synced_at" TIMESTAMP)`

	_, err = db.Exec(cmd)
	if err != nil {
		panic(err)
	}
}

func InsertRecord(article RSS2_Item) {
	db, err := sql.Open("sqlite3", "feed.db")
	if err != nil {
		panic(err)
	}

	cmd := `INSERT OR IGNORE INTO "articles" ("url", "title") VALUES (?, ?)`

	_, err = db.Exec(cmd, article.Link, article.Title)
	if err != nil {
		panic(err)
	}
}
