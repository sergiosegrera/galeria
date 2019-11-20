package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func adminLogin(c *gin.Context) {
	if settings.AdminPassword == "" {
		c.HTML(200, "login.tmpl", gin.H{
			"settings":     settings,
			"pageName":     "create password",
			"inputMessage": "Create a password",
		})
		return
	} else {
		c.HTML(200, "login.tmpl", gin.H{
			"settings":     settings,
			"pageName":     "log in",
			"inputMessage": "Log in",
		})
		return
	}
}

func adminLoginPost(c *gin.Context) {
	password := c.PostForm("password")
	if settings.AdminPassword == "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			log.Println(err.Error())
		}

		settings.AdminPassword = string(hash)
		_, err = database.Exec("UPDATE settings SET value = $2 WHERE name = $1", "AdminPassword", settings.AdminPassword)
		if err != nil {
			log.Println(err.Error())
		}
		c.JSON(200, gin.H{"password": hash})
		return
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(settings.AdminPassword), []byte(password))
		if err != nil {
			c.HTML(200, "login.tmpl", gin.H{
				"settings":      settings,
				"pageName":      "log in",
				"inputMessage":  "Log in",
				"wrongPassword": true,
			})
			return
		}
		c.JSON(200, gin.H{"error": "none"})
	}
}
