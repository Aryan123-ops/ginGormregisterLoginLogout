package main

import (
	"fmt"
	"ginapp/controller"
	"ginapp/database"
	"ginapp/middleware"
	"ginapp/model"

	// "log"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	// loadEnv()
	loadDatabase()
	serveApplication()
}

// func loadEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.POST("/logout", controller.Logout)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
