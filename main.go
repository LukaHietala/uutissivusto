package main

import (
	"uutissivusto/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		articles, err := database.GetArticles(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		categories, err := database.GetCategories(db)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "index.tmpl", gin.H{
			"Articles":   articles,
			"Categories": categories,
		})
	})

	router.GET("/:category", func(c *gin.Context) {
		category := c.Param("category")
		articles, err := database.GetCategoryArticles(db, category)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "category.tmpl", gin.H{
			"Articles": articles,
		})
	})

	router.Run()

}
