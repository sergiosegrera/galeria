package main

import (
	"database/sql"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var database *sql.DB

func main() {
	database, _ = sql.Open("sqlite3", "./galeria.db")
	defer database.Close()

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*/*.tmpl")

	router.GET("/", home)
	router.GET("/about", about)
	router.GET("/admin", admin)

	router.NoRoute(notFound)

	router.Run(":8080")
	log.Println("Started on port 8080")
}
