package main

import (
	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ahmedazizabbassi/pass/api/docs"
	"github.com/ahmedazizabbassi/pass/internal/auth"
	"github.com/ahmedazizabbassi/pass/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"db":     "connected",
		})
	})

	authRouter := router.Group("/auth")
	auth.NewHandler(
		auth.NewService(auth.NewRepository(database.DB)),
	).RegisterRoutes(authRouter)

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
