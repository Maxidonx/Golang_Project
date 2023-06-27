package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var urlMapping = make(map[string]string)

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenURL)
	r.GET("/:shortURL", redirectURL)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func shortenURL(c *gin.Context) {
	longURL := c.PostForm("url")
	shortURL := generateShortURL()

	urlMapping[shortURL] = longURL

	c.JSON(http.StatusOK, gin.H{
		"shortURL": shortURL,
	})
}

func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	longURL, ok := urlMapping[shortURL]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	shortURL := strconv.FormatInt(rand.Int63(), 36)
	return shortURL
}
