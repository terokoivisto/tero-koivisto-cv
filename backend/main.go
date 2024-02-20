package main

import (
	"backend/api"
	"backend/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	dynamo := db.SetupDB()

	r := gin.Default()

	config := cors.Default()
	r.Use(config)

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.CVRoutes(r, dynamo)

	err := r.Run(":7777")
	log.Fatalln(err)
}
