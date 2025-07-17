package domain

import "time"

type Locker struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	LockerNumber uint      `json:"locker_number" gorm:"not null;unique"`
	IsUsed       bool      `json:"is_used" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}