package handler

import (
	"encoding/json"
	"fmt"
	"main-service/internal/usecase"
	"main-service/pkg/util"
	"net/http"
	"time"
	
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

type CreateAttendanceByFingerprintInput struct {
	FingerprintTemplate string `json:"fingerprint_template" binding:"required"`
}

type ClockInInput struct {
    FingerprintID string `json:"fingerprint_id" binding:"required"`
}

type ClockInByFingerprintInput struct {
    FingerprintID string `json:"fingerprint_id" binding:"required"`
}

func (h *AttendanceHandler) ClockInByFingerprint(c *gin.Context) {
    var input ClockInByFingerprintInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Input tidak valid"})
        return
    }

    // Panggil user-service untuk cari user by fingerprint_id
    userServiceURL := fmt.Sprintf("http://user-service:8081/api/v1/users/by-fingerprint/%s", input.FingerprintID)
    resp, err := http.Get(userServiceURL)
    if err != nil || resp.StatusCode != http.StatusOK {
        c.JSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
        return
    }
    defer resp.Body.Close()

    // Decode user dari response user-service
    var userResp struct {
        User map[string]interface{} `json:"user"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal decode user"})
        return
    }
    user := userResp.User
    userID, ok := user["id"].(float64) // JSON number jadi float64
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "User ID tidak valid"})
        return
    }

    // Buat record absensi
    attendance, err := h.attendanceUsecase.CreateAttendance(uint(userID), time.Now().Format(time.RFC3339))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan absensi"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Absensi berhasil",
        "user": user,
        "attendance": attendance,
    })
}


func (h *AttendanceHandler) CreateAttendanceByFingerprint(c *gin.Context) {
	var input CreateAttendanceByFingerprintInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	attendance, err := h.attendanceUsecase.CreateAttendanceByFingerprint(input.FingerprintTemplate)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := util.APIResponse("Attendance created by fingerprint", http.StatusCreated, attendance)
	c.JSON(http.StatusCreated, response)
}
