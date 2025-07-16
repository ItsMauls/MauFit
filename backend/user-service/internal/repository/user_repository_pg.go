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

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
