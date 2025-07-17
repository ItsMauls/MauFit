package repository

import (
	"main-service/internal/domain"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Create(attendance *domain.Attendance) (*domain.Attendance, error) {
	if err := r.db.Create(attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *AttendanceRepository) FindByID(id string) (*domain.Attendance, error) {
	var attendance domain.Attendance
	if err := r.db.First(&attendance, id).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepository) FindAll() ([]*domain.Attendance, error) {
	var attendances []*domain.Attendance
	if err := r.db.Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}
