package main

import (
	routes "go-auth/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(r *gin.Context) {
		r.JSON(200, gin.H{"Success": "Access granted"})
	})

	router.GET("/api/", func(r *gin.Context) {
		r.JSON(200, gin.H{"Success": "Access granted + hi"})
	})

	router.Run(":" + port)
}
