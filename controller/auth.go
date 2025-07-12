package controller

import (
	"net/http"
	"student-service/config"
	"student-service/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(context *gin.Context){
	var body struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := context.ShouldBindJSON(&body)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid input",
		})
		return 
	}

	if body.Username != "Wang" || body.Password != "123456"{
		context.JSON(http.StatusUnauthorized,gin.H{
			"error":"Invalid Login",
		})
	}

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &middleware.Claims{
		Username: body.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(config.JwtKey)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "token": tokenstr,
    })
}

