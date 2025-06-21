package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/config"
)

func AuthMiddleware() gin.HandlerFunc {
	cnf := config.Get()
	var secretKey = []byte(cnf.Server.SecretKey)

	return func(c *gin.Context) {
		skipPaths := []string{
			"/api/v1/login",
			"/api/v1/register",
		}
	
		// Cek apakah path sekarang termasuk yang dikecualikan
		requestPath := c.Request.URL.Path
		for _, path := range skipPaths {
			if strings.HasPrefix(requestPath, path) {
				c.Next() // lewati auth
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		// log.Printf("token : %v", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Split the token string from the Bearer prefix
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims and add to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Printf("claims: %v", claims)

			if userID, ok := claims["user_id"].(float64); ok {
				c.Set("user_id", int(userID)) // Convert float64 to int
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		c.Next()
	}
}
