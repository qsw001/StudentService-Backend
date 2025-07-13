package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
	"student-service/config"
)

type Claims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        if tokenStr == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }
    
        parts := strings.SplitN(tokenStr, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }
   
        tokenStr = parts[1]
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
            return config.JwtKey, nil
        })
    
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
            c.Abort()
            return
        }
    
        // 保存用户信息到 context 中
        c.Set("username", claims.Username)
        c.Next()
    }
}