package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

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
