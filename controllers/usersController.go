package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		return // Stops if error
	}

	// Hasing the password with bcrypt, using the 10 as cost (default)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return // Stops if error
	}

	// Create the user
	// Takes the email from the body and the hashed password
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user",
		})
		return // Stops if error
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// Get the email and password from request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return // Stops if error
	}
	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email and/or password",
		})
		return // Stops if error
	}

	// Compare password in request with stored user password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email and/or password",
		})
		return // Stops if error
	}

	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                                    // Subject is the User id
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expires in 30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) // Load env secret for token

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create JWT",
		})
		return // Stops if error
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) // Sending a 30day valid cookie

	// Send a 200
	c.JSON(http.StatusOK, gin.H{
		// To send JWT -> "token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user") // fetch user data

	c.JSON(http.StatusOK, gin.H{
		"message": user, // Displays the user data
	})
}
