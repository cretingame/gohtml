package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("gohtml server draft")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", nil)
	})

	router.GET("/youtubefeed.html", func(ctx *gin.Context) {
		channelId := "UC_yP2DpIgs5Y1uWC0T03Chw"
		feed, err := getYoutubeFeed(channelId)
		if err != nil {
			ctx.Status(http.StatusTeapot)
			return
		}
		ctx.HTML(http.StatusOK, "youtube.feed.tmpl", gin.H{
			"title": feed.Title,
			"items": feed.Items,
		})
	})

	router.GET("/youtube/channel/:id", func(ctx *gin.Context) {
		var channel struct {
			Id string `uri:"id" binding:"required"`
		}
		err := ctx.ShouldBindUri(&channel)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		feed, err := getYoutubeFeed(channel.Id)
		if err != nil {
			ctx.String(http.StatusNotFound, "Channel %s not found, %s", channel.Id, err.Error())
			return
		}
		ctx.HTML(http.StatusOK, "youtube.feed.tmpl", gin.H{
			"title": feed.Title,
			"items": feed.Items,
		})
	})

	router.GET("/youtube/video/:id", func(ctx *gin.Context) {
		var video struct {
			Id string `uri:"id" binding:"required"`
		}
		err := ctx.ShouldBindUri(&video)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.HTML(http.StatusOK, "youtube.video.tmpl", gin.H{
			"title": "TBD",
			"id":    video.Id,
		})
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

	router.GET("/uri/:id", func(ctx *gin.Context) {
		var channel struct {
			Id string `uri:"id" binding:"required"`
		}
		err := ctx.ShouldBindUri(&channel)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.String(http.StatusOK, "Hello %s !", channel.Id)
	})

	router.Run(":8080")
}
