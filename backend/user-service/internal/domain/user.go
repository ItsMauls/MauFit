package domain

import "time"

type User struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	Name          string       `json:"name" gorm:"not null"`
	Email         string       `json:"email" gorm:"unique;not null"`
	Password      string       `json:"-" gorm:"not null"`
	Role          string       `json:"role" gorm:"type:varchar(16);default:'member';not null"`
	FingerprintID string       `json:"fingerprint_id" gorm:"unique"`
	UserProfile   *UserProfile `json:"user_profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type UserProfile struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null;uniqueIndex"`
	PhotoProfileURL string    `json:"photo_profile_url"`
	Address         string    `json:"address"`
	Phone           string    `json:"phone"`
	Bio             string    `json:"bio"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	User            *User     `json:"-" gorm:"foreignKey:UserID"`
}
