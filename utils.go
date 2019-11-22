package main

import (
	"github.com/gin-gonic/gin"
)

func message(c *gin.Context, code int, n string, m string) {
	c.HTML(code, "message.tmpl", gin.H{
		"settings": settings,
		"pageName": n,
		"message":  m,
	})
}
