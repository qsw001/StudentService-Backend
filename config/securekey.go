package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtKey []byte

func InitSecurekey() {
	err := godotenv.Load()
	if err != nil {
		log.Println("未找到 .env 文件，使用默认配置")
	}

	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		secret = "default_insecure_key" 
	}
	JwtKey = []byte(secret)
}
