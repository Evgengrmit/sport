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

func (a *AuthCodeService) UpdateCode(code authRepo.CodeStatus) error {
	return a.repo.UpdateCode(code)
}
func (a *AuthCodeService) GetCodeStatusByID(code authRepo.AuthorizationCode) (authRepo.CodeStatus, error) {
	return a.repo.GetCodeStatusByID(code)
}
func (a *AuthCodeService) IsCodeExists(id int) (bool, error) {
	return a.repo.IsCodeExists(id)
}
