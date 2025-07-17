package repository

import (
	"main-service/internal/domain"

	"gorm.io/gorm"
)

type LockerRepository struct {
	db *gorm.DB
}

func NewLockerRepository(db *gorm.DB) *LockerRepository {
	return &LockerRepository{db: db}
}

func (r *LockerRepository) Create(locker *domain.Locker) (*domain.Locker, error) {
	if err := r.db.Create(locker).Error; err != nil {
		return nil, err
	}
	return locker, nil
}

func (r *LockerRepository) FindByID(id string) (*domain.Locker, error) {
	var locker domain.Locker
	if err := r.db.First(&locker, id).Error; err != nil {
		return nil, err
	}
	return &locker, nil
}

func (r *LockerRepository) FindAll() ([]*domain.Locker, error) {
	var lockers []*domain.Locker
	if err := r.db.Find(&lockers).Error; err != nil {
		return nil, err
	}
	return lockers, nil
}
