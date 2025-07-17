package usecase

import (
	"main-service/internal/domain"
	"time"
)

type AttendanceUsecase interface {
	CreateAttendance(userID uint, timeIn string) (*domain.Attendance, error)
	GetAttendanceByID(id string) (*domain.Attendance, error)
	GetAllAttendances() ([]*domain.Attendance, error)
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
