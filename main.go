package main

import (
	"github.com/gin-gonic/gin"
	"github.com/twk-mn/goJWT/initializers"
)

func init() {
	initializers.LoadEnvVariables() // Loading .env
	initializers.ConnectToDB()      // Connect to the DB
	initializers.SyncDatabase()     //	Makes sure DB is set up
}

func main() {

	r := gin.Default() // Setting up Gin

	r.GET("/test", func(c *gin.Context) { // test route
		c.JSON(200, gin.H{
			"message": "Yay!",
		})
	})

	r.Run()
}
