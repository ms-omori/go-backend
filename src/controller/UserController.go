package controller

import (
	"ms-omori/domain"
	"ms-omori/usecase"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(uuc usecase.UserUseCase) UserController {
	return UserController{UserUseCase: uuc}
}

func (uc *UserController) GetUserFromName(name string) (*domain.User, error) {
	return uc.UserUseCase.GetUserFromName(name)
}

func (uc *UserController) ListUsers() (*[]domain.User, error) {
	return uc.UserUseCase.ListUsers()
}
