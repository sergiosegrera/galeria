package main

import (
	"database/sql"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	settings Settings
	database *sql.DB
)

func main() {
	port := os.Getenv("PORT")
	databaseUrl, _ := os.LookupEnv("DATABASE_URL")

	if port == "" {
		port = "8080"
	}

	database, _ = sql.Open("postgres", databaseUrl)
	defer database.Close()

	loadSettings()

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*/*.tmpl")

	router.GET("/", home)
	router.GET("/about", about)

	router.GET("/admin", adminLogin)
	router.POST("/admin", adminLoginPost)

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
		case "AdminPassword":
			settings.AdminPassword = value
		}
	}
	rows.Close()

	if settings.WebsiteName == "" {
		_, err := database.Exec("INSERT INTO settings (name, value) VALUES ($1, $2)", "WebsiteName", "Galeria")
		if err != nil {
			log.Println(err.Error())
		}
		settings.WebsiteName = "Galeria"
	}

	if settings.AdminPassword == "" {
		_, err := database.Exec("INSERT INTO settings (name, value) VALUES ($1, $2)", "AdminPassword", "")
		if err != nil {
			log.Println(err.Error())
		}
		settings.AdminPassword = ""
	}
}
