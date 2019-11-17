package main

import (
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(200, "home.tmpl", gin.H{
		"pageName":    "home",
		"websiteName": "galeria",
	})
}

func about(c *gin.Context) {
	return
}

func notFound(c *gin.Context) {
	return
}
