package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("gohtml server draft")

	feedUrl := "https://www.youtube.com/feeds/videos.xml?channel_id=UC_yP2DpIgs5Y1uWC0T03Chw"
	feed, err := getYoutubeFeed(feedUrl)
	if err != nil {
		panic(err)
	}
	items := feed.Items

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.head.tmpl", nil)
		ctx.HTML(http.StatusOK, "youtubeFeed.items.tmpl", gin.H{
			"items": items,
		})
		ctx.HTML(http.StatusOK, "index.foot.tmpl", nil)
	})

	router.GET("/youtubefeed.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.head.tmpl", nil)
		ctx.HTML(http.StatusOK, "youtubeFeed.items.tmpl", gin.H{
			"items": items,
		})
		ctx.HTML(http.StatusOK, "index.foot.tmpl", nil)
	})

	router.GET("/page.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "page.tmpl", gin.H{
			"title":   "Page Title",
			"content": "Page Content",
		})
	})

	router.GET("/string.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello !")
	})

	router.GET("/write.html", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("TEST"))
	})

	router.Run(":8080")
}
