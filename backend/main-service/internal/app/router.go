package app

import (
	"main-service/internal/app/handler"
	"main-service/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, attendanceHandler *handler.AttendanceHandler, lockerHandler *handler.LockerHandler) {
	// Public API routes (no authentication required)
	publicApiV1 := router.Group("/api/v1")
	{
		// Fingerprint attendance routes - no auth required
		publicApiV1.POST("/attendances/fingerprint", attendanceHandler.CreateAttendanceByFingerprint)
		publicApiV1.POST("/attendances/clock-in", attendanceHandler.ClockInByFingerprint)
	}

	// Protected API routes (authentication required)
	protectedApiV1 := router.Group("/api/v1")
	protectedApiV1.Use(middleware.AuthMiddleware())
	{
		attendanceRoutes := protectedApiV1.Group("/attendances")
		{
			attendanceRoutes.POST("", attendanceHandler.CreateAttendance)
			attendanceRoutes.GET("/:id", attendanceHandler.GetAttendanceByID)
			attendanceRoutes.GET("", attendanceHandler.GetAllAttendances)
		}
		lockerRoutes := protectedApiV1.Group("/lockers")
		{
			lockerRoutes.POST("", lockerHandler.CreateLocker)
			lockerRoutes.GET("/:id", lockerHandler.GetLockerByID)
			lockerRoutes.GET("", lockerHandler.GetAllLockers)
		}
	}
}
