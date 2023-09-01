package main

import (
	"encoding/json"
	"fmt"

	"github.com/mmcdole/gofeed"
)

type YoutubeFeed struct {
	Title string
	Items []YoutubeItem
}

type YoutubeItem struct {
	Id          int
	Title       string
	Author      string
	Description string
	Link        string
	Extensions  string
}

func getJDGfeed() ([]*gofeed.Item, error) {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL("https://www.youtube.com/feeds/videos.xml?channel_id=UC_yP2DpIgs5Y1uWC0T03Chw")
	if err != nil {
		return nil, err
	}
	fmt.Println("Feed Title:", feed.Title)

	for i, feedItem := range feed.Items {
		fmt.Println("# Item number", i)
		fmt.Printf("%+v\n", feedItem)
	}
	return feed.Items, err
}

func getYoutubeFeed(url string) (*YoutubeFeed, error) {
	feedParser := gofeed.NewParser()
	rssFeed, err := feedParser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	ytFeed := &YoutubeFeed{
		Title: rssFeed.Title,
	}
	fmt.Println("Feed Title:", rssFeed.Title)

	for i, feedItem := range rssFeed.Items {
		extensionsBytes, err := json.MarshalIndent(feedItem.Extensions, "", "  ")
		if err != nil {
			return nil, err
		}
		ytItem := YoutubeItem{
			Id:          i,
			Title:       feedItem.Title,
			Author:      feedItem.Author.Name,
			Description: feedItem.Extensions["media"]["group"][0].Children["description"][0].Value,
			Link:        feedItem.Link,
			Extensions:  string(extensionsBytes),
		}
		fmt.Println("# Item number", i)
		fmt.Printf("%+v\n", feedItem)
		ytFeed.Items = append(ytFeed.Items, ytItem)
	}
	return ytFeed, err
}
