package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWar() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("account")
		if err == nil {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "未登录",
		})
		c.Abort()
		return
	}
}
