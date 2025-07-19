package app

import (
	"main-service/internal/app/handler"
	"main-service/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, attendanceHandler *handler.AttendanceHandler, lockerHandler *handler.LockerHandler) {
	router.Use(middleware.AuthMiddleware())

	apiV1 := router.Group("/api/v1")
	{
		attendanceRoutes := apiV1.Group("/attendances")
		{
			attendanceRoutes.POST("", attendanceHandler.CreateAttendance)
			attendanceRoutes.GET("/:id", attendanceHandler.GetAttendanceByID)
			attendanceRoutes.GET("", attendanceHandler.GetAllAttendances)
			attendanceRoutes.POST("/fingerprint", attendanceHandler.CreateAttendanceByFingerprint)
		}
		lockerRoutes := apiV1.Group("/lockers")
		{
			lockerRoutes.POST("", lockerHandler.CreateLocker)
			lockerRoutes.GET("/:id", lockerHandler.GetLockerByID)
			lockerRoutes.GET("", lockerHandler.GetAllLockers)
		}
	}
}
