package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type XML struct {
	Bookmarks []struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
		Date  string `xml:"date"`
		Count int    `xml:"bookmarkcount"`
	} `xml:"item"`
}
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

func main() {
	data := httpGet("https://www.tokyomotion.net/rss")

	fmt.Printf(data)

	fmt.Print("------------------------------------")

	result := rss2_0Feed{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// for _, bookmark := range result.Bookmarks {
	// 	datetime, _ := time.Parse(time.RFC3339, bookmark.Date)

	// 	fmt.Printf("%v\n", datetime.Format("2006/01/02 15:04:05"))
	// 	fmt.Printf("%s - %duser\n", bookmark.Title, bookmark.Count)
	// 	fmt.Printf("%v\n", bookmark.Link)
	// 	fmt.Println()
	// }
}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}
