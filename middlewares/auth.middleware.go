package middlewares

import (
	"net/http"
	"strings"

	"github.com/dev-saiful/umanagement/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}
		// Check if it's a Bearer token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}
		// Validate the token using our utility function
		payload, err := utils.VerifyJWT(bearerToken[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("email", payload.Email)
		ctx.Set("role",payload.Role)
		ctx.Next()
	}
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        role, exists := ctx.Get("role")
        if !exists || role != requiredRole {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            ctx.Abort()
            return
        }
        ctx.Next()
    }
}