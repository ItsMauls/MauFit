package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main-service/internal/domain"
	"net/http"
	"time"
)

// User struct to represent user data from user-service
type User struct {
	ID uint `json:"id"`
}

type AttendanceUsecase interface {
	CreateAttendance(userID uint, timeIn string) (*domain.Attendance, error)
	GetAttendanceByID(id string) (*domain.Attendance, error)
	GetAllAttendances() ([]*domain.Attendance, error)
	CreateAttendanceByFingerprint(fingerprintTemplate string) (*domain.Attendance, error)
}

type attendanceUsecase struct {
	repo AttendanceRepository
}

type AttendanceRepository interface {
	Create(attendance *domain.Attendance) (*domain.Attendance, error)
	FindByID(id string) (*domain.Attendance, error)
	FindAll() ([]*domain.Attendance, error)
}

func NewAttendanceUsecase(repo AttendanceRepository) AttendanceUsecase {
	return &attendanceUsecase{repo: repo}
}

func (uc *attendanceUsecase) CreateAttendance(userID uint, timeIn string) (*domain.Attendance, error) {
	timeInParsed, err := time.Parse(time.RFC3339, timeIn)
	if err != nil {
		return nil, err
	}
	attendance := &domain.Attendance{
		UserID: userID,
		TimeIn: timeInParsed,
	}
	return uc.repo.Create(attendance)
}

func (uc *attendanceUsecase) GetAttendanceByID(id string) (*domain.Attendance, error) {
	return uc.repo.FindByID(id)
}

func (uc *attendanceUsecase) GetAllAttendances() ([]*domain.Attendance, error) {
	return uc.repo.FindAll()
}

func (uc *attendanceUsecase) CreateAttendanceByFingerprint(fingerprintTemplate string) (*domain.Attendance, error) {
	// 1. Panggil user-service untuk verifikasi sidik jari
	userServiceURL := "http://user-service:8080/api/v1/users/verify-fingerprint"
	reqBody, _ := json.Marshal(map[string]string{
		"fingerprint_template": fingerprintTemplate,
	})

	resp, err := http.Post(userServiceURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to call user-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user not found for the given fingerprint")
	}

	// 2. Decode response untuk mendapatkan UserID
	var serviceResponse struct {
		Data User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&serviceResponse); err != nil {
		return nil, fmt.Errorf("failed to decode user-service response: %w", err)
	}

	userID := serviceResponse.Data.ID
	if userID == 0 {
		return nil, fmt.Errorf("invalid user ID from user-service")
	}

	// 3. Buat data absensi
	attendance := &domain.Attendance{
		UserID: userID,
		TimeIn: time.Now(),
	}
	return uc.repo.Create(attendance)
}
