package feed2sqlite

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestDBCreate(t *testing.T) {
	CreateTables()
}

func TestDBInsert(t *testing.T) {
	name := filepath.Join("testdata", "rss2.0.xml")
	bytes, err := ioutil.ReadFile(name)

	if err != nil {
		t.Fatalf("Reading %s: %v", name, err)
	}

	data := string(bytes)
	feed, err := ParseRSS(data)

	InsertRecord(feed.Items[0])
}
