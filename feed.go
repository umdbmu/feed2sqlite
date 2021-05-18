package feed2sqlite

import (
	"encoding/xml"
	"fmt"
)

type (
	RSS2_Feed struct {
		XMLName xml.Name      `xml:"rss"`
		Channel *RSS2_Channel `xml:"channel"`
	}

	RSS2_Channel struct {
		Title       string      `xml:"title"`
		Link        string      `xml:"link"`
		Description string      `xml:"description"`
		Copyright   string      `xml:"copyright"`
		Items       []RSS2_Item `xml:"item"`
	}
	RSS2_Item struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
	}
)

func ParseRSS(body string) (RSS2_Channel, error) {
	result := RSS2_Feed{}
	err := xml.Unmarshal([]byte(body), &result)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return *result.Channel, err
}
