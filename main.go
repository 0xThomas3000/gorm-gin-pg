package main

import (
	"log"
	"net/http"

	"github.com/0xThomas3000/gorm-gin-pg/initializers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

// 1. Load the environment variables with Viper, create a connection pool to the Postgres database
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	// 2. Create a Gin router, assign it to the "server" variable
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// 3. A new router group: to group all the routes that have common middlewares / or the same path prefix.
	router := server.Group("/api")

	// 4. Defined a GET route to the "/api/healthchecker" endpoint.
	//    => Specify the endpoint + a handler.
	//		+ The endpoint: the path to a resource on the server.
	//		+ The handler: the function that will be evoked to perform some business logic or return data to the client.
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
