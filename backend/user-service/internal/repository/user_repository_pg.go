package repository

import (
	"user-service/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Save(user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	// Preload UserProfile after creation
	var createdUser domain.User
	err = r.db.Preload("UserProfile").First(&createdUser, user.ID).Error
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("UserProfile").First(&user, id).Error
	if err != nil {
		return nil, err // GORM akan mengembalikan gorm.ErrRecordNotFound jika tidak ada
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Preload("UserProfile").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUserProfile(profile *domain.UserProfile) (*domain.UserProfile, error) {
	var existing domain.UserProfile
	err := r.db.Where("user_id = ?", profile.UserID).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new profile if not exists
			err = r.db.Create(profile).Error
			if err != nil {
				return nil, err
			}
			return profile, nil
		}
		return nil, err
	}
	// Update fields
	existing.Address = profile.Address
	existing.Phone = profile.Phone
	existing.Bio = profile.Bio
	existing.PhotoProfileURL = profile.PhotoProfileURL
	err = r.db.Save(&existing).Error
	if err != nil {
		return nil, err
	}
	return &existing, nil
}

func (r *userRepository) FindByFingerprintID(fingerprintID string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("UserProfile").Where("fingerprint_id = ?", fingerprintID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
