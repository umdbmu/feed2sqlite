package feed2sqlite

import (
	"encoding/xml"
	"time"
)

type (
	rss2_0Feed struct {
		XMLName xml.Name   `xml:"rss"`
		Channel RSSChannel `xml:"channel"`
	}

	RSSChannel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Copyright   string    `xml:"copyright"`
		Items       []RSSItem `xml:"item"`
	}
	RSSItem struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		PubDate     time.Time `xml:"pubDate"`
	}
)

func ParseRSS(body string) (RSSChannel, error) {
	result := rss2_0Feed{}
	err := xml.Unmarshal([]byte(body), &result)

	return result.Channel, err
}
