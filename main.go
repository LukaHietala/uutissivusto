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
	router.Static("/static", "web/static/")
	router.LoadHTMLGlob("web/templates/*")

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

		c.HTML(200, "index.html", gin.H{
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

		categories, err := database.GetCategories(db)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "index.html", gin.H{
			"Articles":   articles,
			"Categories": categories,
		})
	})

	router.GET("/artikkeli/:uri", func(c *gin.Context) {
		articleUri := c.Param("uri")
		article, err := database.GetArticle(db, articleUri)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		categories, err := database.GetCategories(db)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "article.html", gin.H{
			"Article":    article,
			"Categories": categories,
		})

	})

	router.Run()

}
