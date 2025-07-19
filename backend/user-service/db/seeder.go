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

	// Create users with fingerprint_ids that match the simulated users in fingerprint agent
	users := []domain.User{
		{
			Name:          "Admin User",
			Email:         "admin@maufit.com",
			Password:      string(adminPassword),
			Role:          "admin",
			FingerprintID: "ADMIN001", // Special admin fingerprint
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/admin.jpg",
				Address:         "Admin Office, MauFit Gym",
				Phone:           "+621234567890",
				Bio:             "Administrator MauFit Gym",
			},
		},
		{
			Name:          "John Doe",
			Email:         "john.doe@email.com",
			Password:      string(memberPassword),
			Role:          "member",
			FingerprintID: "FP001", // Matches fingerprint agent simulation
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/john.jpg",
				Address:         "Jl. Sudirman No. 123, Jakarta",
				Phone:           "+621234567891",
				Bio:             "Fitness enthusiast, loves cardio workouts",
			},
		},
		{
			Name:          "Jane Smith",
			Email:         "jane.smith@email.com",
			Password:      string(memberPassword),
			Role:          "member",
			FingerprintID: "FP002", // Matches fingerprint agent simulation
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/jane.jpg",
				Address:         "Jl. Thamrin No. 456, Jakarta",
				Phone:           "+621234567892",
				Bio:             "Yoga instructor and wellness coach",
			},
		},
		{
			Name:          "Bob Johnson",
			Email:         "bob.johnson@email.com",
			Password:      string(memberPassword),
			Role:          "member",
			FingerprintID: "FP003", // Matches fingerprint agent simulation
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/bob.jpg",
				Address:         "Jl. Gatot Subroto No. 789, Jakarta",
				Phone:           "+621234567893",
				Bio:             "Weightlifting champion, personal trainer",
			},
		},
		{
			Name:          "Alice Brown",
			Email:         "alice.brown@email.com",
			Password:      string(memberPassword),
			Role:          "member",
			FingerprintID: "FP004", // Matches fingerprint agent simulation
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/alice.jpg",
				Address:         "Jl. Kuningan No. 321, Jakarta",
				Phone:           "+621234567894",
				Bio:             "Swimming coach and aqua fitness trainer",
			},
		},
		{
			Name:          "Charlie Wilson",
			Email:         "charlie.wilson@email.com",
			Password:      string(memberPassword),
			Role:          "member",
			FingerprintID: "FP005", // Matches fingerprint agent simulation
			UserProfile: &domain.UserProfile{
				PhotoProfileURL: "https://example.com/photos/charlie.jpg",
				Address:         "Jl. Senayan No. 654, Jakarta",
				Phone:           "+621234567895",
				Bio:             "CrossFit athlete and nutrition specialist",
			},
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed user data: %v", err)
	}

	log.Println("User data seeding successful.")
}
