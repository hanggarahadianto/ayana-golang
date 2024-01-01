// main.go
package main

import (
	"ayana/db"
	"ayana/routes"
	"ayana/utils"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configure, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}
	db.InitializeDb(&configure)
	gin.SetMode(gin.ReleaseMode)

	// Set up Gin router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my Ayana application! 🚀",
		})
		fmt.Println("Welcome to my Ayana application! 🚀")
	})

	routes.SetupHomeRouter(router)
	routes.SetupReservationRouter(router)

	// defineRoutes(routes)

	fmt.Println("running on server : " + configure.ServerPort)
	log.Fatal(router.Run(":" + configure.ServerPort))
}
