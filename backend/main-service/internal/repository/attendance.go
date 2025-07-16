package repository

import (
	"main-service/internal/domain"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func (r *AttendanceRepository) Save(attendace *domain.Attendance) (*domain.Attendance, error) {
	err := r.db.Create(attendace).Error

	if(err != nil) {
		return nil, err
	}
	var createdAttendance domain.Attendance
	
	return &createdAttendance, nil
}

func (r *AttendanceRepository) FindAll() ([] *domain.Attendance, error) {
	var attendance []*domain.Attendance

	return attendance, nil
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{}
}