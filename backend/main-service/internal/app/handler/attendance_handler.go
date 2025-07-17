package handler

import (
	"encoding/json"
	"fmt"
	"main-service/internal/usecase"
	"main-service/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceUsecase usecase.AttendanceUsecase
}
type AttendanceWithUser struct {
	Attendance interface{} `json:"attendance"`
	User       interface{} `json:"user"`
}

func NewAttendanceHandler(uc usecase.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceUsecase: uc,
	}
}

type CreateAttendanceInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	TimeIn string `json:"time_in" binding:"required"`
}

func (h *AttendanceHandler) CreateAttendance(c *gin.Context) {
	var input CreateAttendanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Call user-service to validate user
	userServiceURL := fmt.Sprintf("http://user-service:8080/api/v1/users/%d", input.UserID)
	resp, err := http.Get(userServiceURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		response := util.APIResponse("User not found in user-service", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	resp.Body.Close()
	attendance, err := h.attendanceUsecase.CreateAttendance(input.UserID, input.TimeIn)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Attendance created", http.StatusCreated, attendance)
	c.JSON(http.StatusCreated, response)
}

func (h *AttendanceHandler) GetAttendanceByID(c *gin.Context) {
	id := c.Param("id")
	attendance, err := h.attendanceUsecase.GetAttendanceByID(id)
	if err != nil {
		response := util.APIResponse("Attendance not found", http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := util.APIResponse("Attendance found", http.StatusOK, attendance)
	c.JSON(http.StatusOK, response)
}

func (h *AttendanceHandler) GetAllAttendances(c *gin.Context) {
	attendances, err := h.attendanceUsecase.GetAllAttendances()
	if err != nil {
		response := util.APIResponse("Failed to get attendances", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var results []AttendanceWithUser
	for _, attendance := range attendances {
		// Pastikan attendance punya field UserID
		userID := attendance.UserID // sesuaikan dengan struct attendance kamu

		userServiceURL := fmt.Sprintf("http://user-service:8080/api/v1/users/%d", userID)
		resp, err := http.Get(userServiceURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			// Bisa skip user, atau masukkan nil, atau return error, sesuai kebutuhan
			results = append(results, AttendanceWithUser{
				Attendance: attendance,
				User:       nil,
			})
			if resp != nil {
				resp.Body.Close()
			}
			continue
		}
		var user interface{}
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			user = nil
		}
		resp.Body.Close()

		results = append(results, AttendanceWithUser{
			Attendance: attendance,
			User:       user,
		})
	}

	response := util.APIResponse("Attendances with users found", http.StatusOK, results)
	c.JSON(http.StatusOK, response)
}
