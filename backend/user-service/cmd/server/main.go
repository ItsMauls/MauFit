package main

import (
	"log"
	"user-service/configs"
	"user-service/db"
	"user-service/internal/app"
	"user-service/internal/app/handler"
	"user-service/internal/domain"
	"user-service/internal/repository"
	"user-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	// 1. Initialize Database
	database, err := configs.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 2. Run Automatic Migration
	log.Println("Running database migration...")
	err = database.AutoMigrate(&domain.UserProfile{}, &domain.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration successful.")

	// 3. Run Seeding
	db.SeedUsers(database) // NEW: Call the seeder function

	// 4. Initialize Application Layers
	userRepo := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// 5. Setup & Run Server
	router := gin.Default()
	app.SetupRouter(router, userHandler)

	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
