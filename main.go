package main

import (
	"database/sql"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	settings Settings
	database *sql.DB
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	database, _ = sql.Open("sqlite3", "./galeria.db")
	defer database.Close()

	loadSettings()

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*/*.tmpl")

	router.GET("/", home)
	router.GET("/about", about)
	router.GET("/admin", admin)

	router.NoRoute(notFound)

	router.Run(":" + port)
	log.Println("Started on port 8080")
}

func loadSettings() {
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS settings (name TEXT, value TEXT)")
	statement.Exec()

	rows, _ := database.Query("SELECT name, value FROM settings")
	var name string
	var value string
	for rows.Next() {
		rows.Scan(&name, &value)
		switch name {
		case "WebsiteName":
			settings.WebsiteName = value
		}
	}
	rows.Close()
}
