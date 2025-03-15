package main

import (
	_ "github.com/go-playground/validator/v10"
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/gofiber/storage/redis"
	_ "github.com/joho/godotenv"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/sirupsen/logrus"
	_ "github.com/stretchr/testify"
	_ "golang.org/x/crypto/bcrypt"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

func main() {}
