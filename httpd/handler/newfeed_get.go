package handler

import (
	"net/http"
	"newfeed/flatform/newfeed"

	"github.com/gin-gonic/gin"
)

func NewFeedGet(feed *newfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
