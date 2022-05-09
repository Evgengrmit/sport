package authService

import "sport/pkg/repository/authRepo"

type Authorization interface {
	CreateUser(u authRepo.User) (authRepo.User, error)
}

type AuthorizationService struct {
	repo authRepo.AuthorizationRepo
}
