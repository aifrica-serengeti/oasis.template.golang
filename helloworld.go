package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Define the handler function
	router.GET("/", func(c *gin.Context) {
		log.Print("helloworld: received a request")
		target := os.Getenv("TARGET")
		if target == "" {
			target = "World"
		}
		c.String(200, fmt.Sprintf("Hello %s!\n", target))
	})

	// Determine the port to listen on
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)

	// Start the Gin server
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}

