package main

import (
	"github.com/gin-gonic/gin"
	"github.com/twk-mn/goJWT/initializers"
)

func init() {
	initializers.LoadEnvVariables() // Loading .env
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
