package main

import (
	"log"
	"main-service/configs"
	"main-service/db" // tambahkan import db
	"main-service/internal/app"
	"main-service/internal/app/handler"
	"main-service/internal/domain"
	"main-service/internal/repository"
	"main-service/internal/usecase"

	"github.com/gin-contrib/cors"
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

	db.SeedData(database) // panggil seeder locker & attendance

	attendanceRepo := repository.NewAttendanceRepository(database)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepo)
	attendanceHandler := handler.NewAttendanceHandler(attendanceUsecase)

	lockerRepo := repository.NewLockerRepository(database)
	lockerUsecase := usecase.NewLockerUsecase(lockerRepo)
	lockerHandler := handler.NewLockerHandler(lockerUsecase)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	app.SetupRouter(router, attendanceHandler, lockerHandler)

	log.Println("Server is running on http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
