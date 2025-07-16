package domain

import "time"

type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	TimeIn    uint      `json:"time_in"`
	TimeOut   uint      `json:"time_out"`
	LockerID  *uint     `json:"locker_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Locker struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	LockerNumber uint      `json:"locker_number" gorm:"not null;unique"`
	IsUsed       bool      `json:"is_used" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
