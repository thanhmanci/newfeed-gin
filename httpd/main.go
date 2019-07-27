package main

import (
	"newfeed/flatform/newfeed"
	"newfeed/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newfeed.New()

	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/get_feeds", handler.NewFeedGet(feed))
	r.POST("/new_feeds", handler.NewFeedPost(feed))

	r.Run(":3000")
}
