package middlewares

import (
	"apiv2/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// checks if user is authenticated with a jwt token
// AbortWithStatusJSON cancels any task
func Authenticated(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	// Check if the Authorization header is missing
	if tokenHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}
	// Ensure the token follows the "Bearer <token>" format
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(tokenHeader, bearerPrefix) {
		// stop if error
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	// Extract the actual token by removing the "Bearer " prefix
	token := strings.TrimPrefix(tokenHeader, bearerPrefix)

	// Verify the token
	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(c.ClientIP())
	// attach data
	c.Set("userID", userID)
	// next line continues or pending
	c.Next()
}
