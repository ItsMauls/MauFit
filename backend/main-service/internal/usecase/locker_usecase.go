package usecase

import (
	"main-service/internal/domain"
)

type LockerUsecase interface {
	CreateLocker(number uint, status bool) (*domain.Locker, error)
	GetLockerByID(id string) (*domain.Locker, error)
	GetAllLockers() ([]*domain.Locker, error)
}

type lockerUsecase struct {
	repo LockerRepository
}

type LockerRepository interface {
	Create(locker *domain.Locker) (*domain.Locker, error)
	FindByID(id string) (*domain.Locker, error)
	FindAll() ([]*domain.Locker, error)
}

func NewLockerUsecase(repo LockerRepository) LockerUsecase {
	return &lockerUsecase{repo: repo}
}

func (uc *lockerUsecase) CreateLocker(number uint, status bool) (*domain.Locker, error) {
	locker := &domain.Locker{
		LockerNumber: number,
		IsUsed: status,
	}
	return uc.repo.Create(locker)
}

func (uc *lockerUsecase) GetLockerByID(id string) (*domain.Locker, error) {
	return uc.repo.FindByID(id)
}

func (uc *lockerUsecase) GetAllLockers() ([]*domain.Locker, error) {
	return uc.repo.FindAll()
}
