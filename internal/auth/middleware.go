package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userkey = "user"

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)

	if user := session.Get(userkey); user == nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}
