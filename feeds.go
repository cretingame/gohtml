package main

import (
	"encoding/json"

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
	VideoId     string
	Extensions  string
}

func getYoutubeFeed(channelId string) (*YoutubeFeed, error) {
	feedParser := gofeed.NewParser()
	url := "https://www.youtube.com/feeds/videos.xml?channel_id=" + channelId
	rssFeed, err := feedParser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	ytFeed := &YoutubeFeed{
		Title: rssFeed.Title,
	}

	for i, feedItem := range rssFeed.Items {
		extensionsBytes, err := json.MarshalIndent(feedItem.Extensions, "", "  ")
		if err != nil {
			return nil, err
		}
		videoId := feedItem.Extensions["yt"]["videoId"][0].Value
		ytItem := YoutubeItem{
			Id:          i,
			Title:       feedItem.Title,
			Author:      feedItem.Author.Name,
			Description: feedItem.Extensions["media"]["group"][0].Children["description"][0].Value,
			// Link:        feedItem.Link,
			Link:       "../video/" + videoId,
			VideoId:    videoId,
			Extensions: string(extensionsBytes),
		}
		ytFeed.Items = append(ytFeed.Items, ytItem)
	}
	return ytFeed, err
}
