package main

import (
	"github.com/gin-gonic/gin"
	"github.com/twk-mn/goJWT/controllers"
	"github.com/twk-mn/goJWT/initializers"
)

func init() {
	initializers.LoadEnvVariables() // Loading .env
	initializers.ConnectToDB()      // Connect to the DB
	initializers.SyncDatabase()     //	Makes sure DB is set up
}

func main() {

	r := gin.Default() // Setting up Gin
	// Route for signup
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run()
}
