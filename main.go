package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zachkp/goshort/handler"
	"github.com/zachkp/goshort/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey GoShort API!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store init happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server - ERROR: %v", err))
	}
}
