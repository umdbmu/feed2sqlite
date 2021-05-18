package feed2sqlite

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestParseItem(t *testing.T) {
	name := filepath.Join("testdata", "rss2.0.xml")
	bytes, err := ioutil.ReadFile(name)

	if err != nil {
		t.Fatalf("Reading %s: %v", name, err)
	}

	data := string(bytes)
	feed, err := ParseRSS(data)
	if err != nil {
		t.Fatalf("Parsing %s: %v", name, err)
	}

	if feed.Title != "RSS Test" {
		t.Fatalf("Parsing Title is :%s", feed.Title)
	}
}
