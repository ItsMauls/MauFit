package main

import (
	"log"
	"main-service/configs"
	"main-service/internal/app"
	"main-service/internal/app/handler"
	"main-service/internal/domain"
	"main-service/internal/repository"
	"main-service/internal/usecase"

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

	attendanceRepo := repository.NewAttendanceRepository(database)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepo)
	attendanceHandler := handler.NewAttendanceHandler(attendanceUsecase)

	lockerRepo := repository.NewLockerRepository(database)
	lockerUsecase := usecase.NewLockerUsecase(lockerRepo)
	lockerHandler := handler.NewLockerHandler(lockerUsecase)

	router := gin.Default()
	app.SetupRouter(router, attendanceHandler, lockerHandler)

	log.Println("Server is running on http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
