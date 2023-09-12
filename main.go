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

	// Canonical ids
	channels := []string{
		"UC_yP2DpIgs5Y1uWC0T03Chw", // Joueur du grenier
		"UCOJjZIqV9ZzFDgnKOtF-NZg", // Léonie K.
	}
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"channels": channels,
		})
	})

	array := [][3]string{
		{"Alfreds Futterkiste", "Maria Anders", "Germany"},
		{"Centro comercial Moctezuma", "Francisco Chang", "Mexico"},
		{"Trader Joes", "James Better", "U.S.A."},
		{"Migros", "Adolf Oggi", "Switzerland"},
		{"Carrefour", "Emmanuel Macron", "France"},
	}
	router.GET("/table.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "table.tmpl", gin.H{
			"array": array,
		})
	})

	router.PUT("/table/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusNotImplemented, "Not implemented")
	})

	router.DELETE("/table/:id", func(ctx *gin.Context) {
		var tableRow struct {
			Id int `uri:"id" binding:"required"`
		}
		err := ctx.ShouldBindUri(&tableRow)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Bind URI error: "+err.Error())
			return
		}

		if tableRow.Id < 0 || tableRow.Id >= len(array) {
			ctx.String(http.StatusBadRequest, "Not in the table: "+err.Error())
			return
		}

		array = append(array[:tableRow.Id], array[tableRow.Id+1:]...)

		ctx.Status(http.StatusOK)
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
