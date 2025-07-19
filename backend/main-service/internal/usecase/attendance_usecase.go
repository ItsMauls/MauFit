package usecase

import (
	"encoding/json"
	"fmt"
	"main-service/internal/domain"
	"net/http"
	"strings"
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
	// 1. Extract fingerprint_id from template using pattern matching
	fingerprintID := uc.extractFingerprintIDFromTemplate(fingerprintTemplate)
	if fingerprintID == "" {
		return nil, fmt.Errorf("could not extract fingerprint_id from template")
	}

	// 2. Panggil user-service untuk cari user by fingerprint_id
	userServiceURL := fmt.Sprintf("http://user-service:8080/api/v1/users/by-fingerprint/%s", fingerprintID)
	resp, err := http.Get(userServiceURL)
	if err != nil {
		return nil, fmt.Errorf("failed to call user-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user not found for fingerprint_id: %s", fingerprintID)
	}

	// 3. Decode response untuk mendapatkan UserID
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

	// 4. Buat data absensi
	attendance := &domain.Attendance{
		UserID: userID,
		TimeIn: time.Now(),
	}
	return uc.repo.Create(attendance)
}

// extractFingerprintIDFromTemplate extracts fingerprint_id from realistic template
func (uc *attendanceUsecase) extractFingerprintIDFromTemplate(template string) string {
	// Extract fingerprint_id from new template format: TEMPLATE_{fingerprint_id}_Q{quality}_{timestamp}
	// Example: TEMPLATE_FP001_Q85_1642678901234
	if strings.HasPrefix(template, "TEMPLATE_") {
		// Remove "TEMPLATE_" prefix
		template = strings.TrimPrefix(template, "TEMPLATE_")

		// Split by underscore and get the first part (fingerprint_id)
		parts := strings.Split(template, "_")
		if len(parts) >= 1 {
			return parts[0] // This should be the fingerprint_id (e.g., "FP001")
		}
	}

	return ""
}
