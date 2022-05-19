package authService

import (
	"sport/pkg/repository/authRepo"
)

func NewAuthorizationService(repo authRepo.AuthorizationRepo) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (a *AuthorizationService) CreateUser(u authRepo.User) (authRepo.AuthorizationCode, error) {
	return a.repo.CreateUser(u)
}

func NewAuthCodeService(repo authRepo.AuthCodeRepo) *AuthCodeService {
	return &AuthCodeService{repo: repo}
}

func (a *AuthCodeService) CreateCode(user authRepo.User) (authRepo.AuthorizationCode, error) {
	return a.repo.CreateCode(user)
}
func (a *AuthCodeService) VerifyCode(code authRepo.AuthorizationCode) (authRepo.User, error) {
	return a.repo.VerifyCode(code)
}
