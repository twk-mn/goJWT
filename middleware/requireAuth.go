package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/twk-mn/goJWT/initializers"
	"github.com/twk-mn/goJWT/models"
)

func RequireAuth(c *gin.Context) {
	// Get cookie from request
	tokenString, err := c.Cookie("Authorization") // fetches the Authorization data from the cookie
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode and validate cookie
	// Checking signing method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // If token is ok
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with the token
		var user models.User
		initializers.DB.First(&user, claims["sub"]) // Getting user Id

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
