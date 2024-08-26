package usecase

import (
	"ms-omori/domain"
	"ms-omori/infrastructure"
)

type UserUseCase interface {
	GetUserFromName(name string) (*domain.User, error)
	ListUsers() (*[]domain.User, error)
}
type UserUseCaseImpl struct {
	UserRepository infrastructure.UserRepository
}

func NewUserUseCase(ur infrastructure.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{UserRepository: ur}
}

func (uuc *UserUseCaseImpl) GetUserFromName(name string) (*domain.User, error) {
	return uuc.UserRepository.GetUserFromName(name)
}

func (uuc *UserUseCaseImpl) ListUsers() (*[]domain.User, error) {
	return uuc.UserRepository.ListUsers()
}
