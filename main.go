package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/kisu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"kuva": "https://cataas.com/cat/fQCYUQNMnhPBN4jb?position=center",
		})
	})
	router.Run()
}
