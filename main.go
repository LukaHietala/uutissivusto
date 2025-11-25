package main

import (
	"uutissivusto/internal/auth"
	"uutissivusto/internal/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/static", "web/static/")
	router.LoadHTMLGlob("web/templates/*")

	// for auth middleware
	store := cookie.NewStore([]byte("salaisuus"))
	router.Use(sessions.Sessions("auth", store))

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

	router.GET("/kirjaudu", auth.Login)
	router.POST("/kirjaudu", auth.LoginPost)
	router.GET("/logout", auth.Logout)

	// all routes in authorized group require auth
	authorized := router.Group("/admin")

	authorized.Use(auth.AuthRequired)

	authorized.GET("/luo", func(c *gin.Context) {
		categories, err := database.GetCategories(db)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "luo.html", gin.H{
			"Categories": categories,
		})
	})

	authorized.POST("/luo", func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")
		picture := c.PostForm("picture")
		description := c.PostForm("description")
		uri := c.PostForm("uri")
		category_id := c.PostForm("categories")

		if title != "" && content != "" && picture != "" && description != "" && uri != "" {
			err := database.AddArticle(db, title, content, picture, description, uri, category_id)

			if err != nil {
				c.HTML(500, "luo.html", gin.H{"Error": "Artikkelin luominen ei onnistunut"})
				return
			}
		}
		categories, err := database.GetCategories(db)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "luo.html", gin.H{
			"Categories": categories,
		})

	})
	router.Run(":8080")

}
