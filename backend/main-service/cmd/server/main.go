package main

import (
	"log"
	"main-service/configs"
	"main-service/internal/app"
	"main-service/internal/app/handler"
	"main-service/internal/domain"
	"main-service/internal/repository"
	"main-service/internal/usecase"
	"user-service/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	database, err := configs.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to initialize Database: %v", err)
	}

	log.Println("Running Database Migration...")
	err = database.AutoMigrate(&domain.Attendance{}, &domain.Locker{})

	if err != nil {
		log.Fatalf("Failed to Migrate Database %v", err)
	}
	log.Println("Database Migration Successful")

	db.seedUsers(database)

	mainRepo := repository.NewMainRepository(database)
	mainUseCase := usecase.NewMainUsecase(mainRepo)
	mainHandler := handler.NewMainHandler(mainUseCase)

	router := gin.Default()
	app.SetupRouter(router, mainHandler)

	log.Println("Server is running on http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
