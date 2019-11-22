package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		session := sessions.Default(c)
		auth := session.Get("authenticated")
		if auth == nil {
			c.HTML(200, "login.tmpl", gin.H{
				"settings":     settings,
				"pageName":     "log in",
				"inputMessage": "Log in",
			})
			return
		} else if auth.(bool) {
			c.HTML(200, "dashboard.tmpl", gin.H{
				"settings": settings,
				"pageName": "dashboard",
			})
		}
	}
}

func adminLoginPost(c *gin.Context) {
	password := c.PostForm("password")
	if settings.AdminPassword == "" && len(password) > 7 {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			message(c, 500, "error", "Internal server error :(")
			return
		}

		settings.AdminPassword = string(hash)
		_, err = database.Exec("UPDATE settings SET value = $2 WHERE name = $1", "AdminPassword", settings.AdminPassword)
		if err != nil {
			message(c, 500, "error", "Internal server error :(")
			return
		}
		message(c, 200, "password created", "Password Created")
		return
	} else if len(password) < 7 {
		c.HTML(200, "login.tmpl", gin.H{
			"settings":      settings,
			"pageName":      "log in",
			"inputMessage":  "Password too short",
			"wrongPassword": true,
		})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(settings.AdminPassword), []byte(password))
		if err != nil {
			c.HTML(200, "login.tmpl", gin.H{
				"settings":      settings,
				"pageName":      "log in",
				"inputMessage":  "Wrong password",
				"wrongPassword": true,
			})
			return
		}
		session := sessions.Default(c)
		session.Set("authenticated", true)
		session.Save()
		c.HTML(200, "dashboard.tmpl", gin.H{
			"settings": settings,
			"pageName": "dashboard",
		})
	}
}
