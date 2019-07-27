package handler

import (
	"fmt"
	"net/http"
	"newfeed/flatform/newfeed"

	"github.com/gin-gonic/gin"
)

type newfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewFeedPost(feed *newfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody := newfeedPostRequest{}
		c.Bind(&reqBody)

		item := newfeed.Item{
			Title: reqBody.Title,
			Post:  reqBody.Post,
		}
		fmt.Println(item.Post)
		feed.Add(item)

		c.JSON(http.StatusOK, item)
	}
}
