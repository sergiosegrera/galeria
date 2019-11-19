package main

import (
	"github.com/gin-gonic/gin"
)

type Settings struct {
	WebsiteName   string
	AdminPassword string
}

func home(c *gin.Context) {
	c.HTML(200, "home.tmpl", gin.H{
		"settings": settings,
		"pageName": "home",
	})
}

func about(c *gin.Context) {
	return
}

func notFound(c *gin.Context) {
	return
}
