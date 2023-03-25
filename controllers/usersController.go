package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twk-mn/goJWT/initializers"
	"github.com/twk-mn/goJWT/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/password from the request body
	// Setting up the structure
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		// Stops if error
		return
	}

	// Hasing the password with bcrypt, using the 10 as cost (default)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		// Stops if error
		return
	}

	// Create the user
	// Takes the email from the body and the hashed password
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user",
		})

		// Stops if error
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
