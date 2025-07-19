package db

import (
	"log"
	"main-service/internal/domain"
	"time"

	"gorm.io/gorm"
)

// SeedData akan mengisi data awal locker dan attendance jika tabel kosong
func SeedData(db *gorm.DB) {
	// Seeder Locker
	var lockerCount int64
	db.Model(&domain.Locker{}).Count(&lockerCount)
	if lockerCount == 0 {
		lockers := make([]domain.Locker, 100)
		for i := 0; i < 100; i++ {
			lockers[i] = domain.Locker{
				LockerNumber: uint(i + 1),
				IsUsed:       i%3 == 0, // Mix: setiap kelipatan 3 terisi, lainnya available
			}
		}
		for _, locker := range lockers {
			db.FirstOrCreate(&locker, domain.Locker{LockerNumber: locker.LockerNumber})
		}
		log.Println("Locker seeding completed.")
	} else {
		log.Println("Locker seeding skipped, data already exists.")
	}

	// Seeder Attendance (contoh, bisa disesuaikan)
	var attendanceCount int64
	db.Model(&domain.Attendance{}).Count(&attendanceCount)
	if attendanceCount == 0 {
		var lockers []domain.Locker
		db.Find(&lockers)
		if len(lockers) >= 2 {
			attendances := []domain.Attendance{
				{UserID: 1, TimeIn: time.Now(), LockerID: &lockers[0].ID},
				{UserID: 2, TimeIn: time.Now(), LockerID: &lockers[1].ID},
			}
			for _, attendance := range attendances {
				db.Create(&attendance)
			}
			log.Println("Attendance seeding completed.")
		} else {
			log.Println("Attendance seeding skipped, not enough lockers.")
		}
	} else {
		log.Println("Attendance seeding skipped, data already exists.")
	}
}
