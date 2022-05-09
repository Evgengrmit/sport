package authService

import (
	"sport/pkg/repository/authRepo"
)

func NewAuthorizationService(repo authRepo.AuthorizationRepo) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (a *AuthorizationService) CreateUser(u authRepo.User) (authRepo.User, error) {
	return a.repo.CreateUser(u)
}
