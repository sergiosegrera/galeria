package main

import (
	"github.com/gin-gonic/gin"
)

func admin(c *gin.Context) {
	if settings.AdminPassword == "" {
		c.HTML(200, "login.tmpl", gin.H{
			"settings": settings,
			"pageName": "create password",
		})
		return
	}
}

func adminPost(c *gin.Context) {
	return
}
