package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func LoginPost(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if password != "" && username != "" {
		if username != "mirri" || password != "kissa" {
			c.HTML(401, "login.html", gin.H{"Error": "Käyttäjää ei löytynyt"})
			return
		}

		session.Set(userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(500, "login.html", gin.H{"Error": "Ongelmissa"})
			return
		}
		c.Redirect(302, "/")
		return
	}

	c.HTML(200, "login.html", gin.H{})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.Redirect(302, "/")
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(500, gin.H{"error": "session save failed"})
		return
	}
	c.Redirect(302, "/")
}
