// Create this new file: /db/seeder.go
package db

import (
	"log"
	"user-service/internal/domain"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

// SeedUsers will populate initial user data if the table is empty.
func SeedUsers(db *gorm.DB) {
	// Check if there are any users in the database
	var userCount int64
	db.Model(&domain.User{}).Count(&userCount)

	if userCount > 0 {
		log.Println("Seeding not required, user data already exists.")
		return
	}

	log.Println("Starting user data seeding...")

	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	memberPassword, _ := bcrypt.GenerateFromPassword([]byte("member123"), bcrypt.DefaultCost)

	users := []domain.User{
		{
			Name:          "Admin User",
			Email:         "admin@example.com",
			Password:      string(adminPassword),
			Role:          "admin",
			FingerprintID: "admin-fp-001",
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/admin.jpg",
				Address:         "Admin Street 1, City",
				Phone:           "+621234567890",
				Bio:             "Super admin user.",
			},
		},
		{
			Name:          "Test User 1",
			Email:         "test1@example.com",
			Password:      string(memberPassword),
			Role:          "member",
			FingerprintID: "member-fp-001",
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/member1.jpg",
				Address:         "Member Street 2, City",
				Phone:           "+621234567891",
				Bio:             "Regular member user.",
			},
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed user data: %v", err)
	}

	log.Println("User data seeding successful.")
}
