package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid Authorization header"})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verifikasi token ke user-service
		req, err := http.NewRequest("GET", "http://user-service:8080/api/v1/users/verify-token", nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
			c.Abort()
			return
		}
		req.Header.Set("Authorization", "Bearer "+token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		// (Opsional) Ambil data user dari response dan set ke context
		// var user map[string]interface{}
		// json.NewDecoder(resp.Body).Decode(&user)
		// c.Set("user", user)

		c.Next()
	}
}