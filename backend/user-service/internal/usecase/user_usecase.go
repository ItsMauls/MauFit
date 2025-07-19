package usecase

import (
	"errors"
	"fmt"
	"user-service/internal/domain"

	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	UpdateUserProfile(profile *domain.UserProfile) (*domain.UserProfile, error)
	FindByFingerprintID(fingerprintID string) (*domain.User, error)
}

type UserUsecase interface {
	CreateUser(name, email string, fingerprintId string) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	RegisterUser(name, email, password, photoProfileURL, address, phone, bio string) (*domain.User, error)
	AdminLogin(email, password string) (string, error)
	VerifyToken(tokenString string) (*domain.User, error)
	UpdateUserProfile(id, address, phone, bio, photoProfileURL string) (*domain.UserProfile, error)
	VerifyFingerprint(fingerprintTemplate string) (*domain.User, error)
}

type userUsecase struct {
	userRepo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (uc *userUsecase) CreateUser(name, email string, fingerprintId string) (*domain.User, error) {
	user := &domain.User{
		Name:          name,
		Email:         email,
		FingerprintID: fingerprintId,
	}
	return uc.userRepo.Save(user)
}

func (uc *userUsecase) GetUserByID(id string) (*domain.User, error) {
	return uc.userRepo.FindByID(id)
}

func (uc *userUsecase) GetAllUsers() ([]*domain.User, error) {
	return uc.userRepo.FindAll()
}

func (uc *userUsecase) RegisterUser(name, email, password, photoProfileURL, address, phone, bio string) (*domain.User, error) {
	// Check if email already exists
	users, err := uc.userRepo.FindAll()
	if err == nil {
		for _, u := range users {
			if u.Email == email {
				return nil, errors.New("email already registered")
			}
		}
	}
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     "member",
		UserProfile: &domain.UserProfile{
			PhotoProfileURL: photoProfileURL,
			Address:         address,
			Phone:           phone,
			Bio:             bio,
		},
	}
	return uc.userRepo.Save(user)
}

func (uc *userUsecase) AdminLogin(email, password string) (string, error) {
	users, err := uc.userRepo.FindAll()
	if err != nil {
		return "", errors.New("user not found")
	}
	var admin *domain.User
	for _, u := range users {
		if u.Email == email && u.Role == "admin" {
			admin = u
			break
		}
	}
	if admin == nil {
		return "", errors.New("admin user not found or not admin")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}
	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": admin.ID,
		"email":   admin.Email,
		"role":    admin.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret" // fallback default
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (uc *userUsecase) VerifyToken(tokenString string) (*domain.User, error) {
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("user_id not found in token")
	}
	user, err := uc.userRepo.FindByID(fmt.Sprintf("%d", int64(userID)))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) UpdateUserProfile(id, address, phone, bio, photoProfileURL string) (*domain.UserProfile, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user.UserProfile == nil {
		user.UserProfile = &domain.UserProfile{UserID: user.ID}
	}
	if address != "" {
		user.UserProfile.Address = address
	}
	if phone != "" {
		user.UserProfile.Phone = phone
	}
	if bio != "" {
		user.UserProfile.Bio = bio
	}
	if photoProfileURL != "" {
		user.UserProfile.PhotoProfileURL = photoProfileURL
	}
	updated, err := uc.userRepo.UpdateUserProfile(user.UserProfile)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (uc *userUsecase) VerifyFingerprint(fingerprintTemplate string) (*domain.User, error) {
	return uc.userRepo.FindByFingerprintID(fingerprintTemplate)
}
