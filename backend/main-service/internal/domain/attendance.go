package domain

import "time"

type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	TimeIn    time.Time  `json:"time_in"`
	TimeOut   time.Time  `json:"time_out"`
	LockerID  *uint     `json:"locker_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

