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
		ctx.HTML(http.StatusOK, "index.head.tmpl", nil)
		ctx.HTML(http.StatusOK, "index.foot.tmpl", nil)
	})

	router.Run(":8080")

}
