package authService

import (
	"sport/pkg/repository/authRepo"
)

type AuthCode interface {
	CreateCode(user authRepo.User) (authRepo.AuthorizationCode, error)
	VerifyCode(code authRepo.AuthorizationCode) (authRepo.User, error)
	UpdateCode(code authRepo.CodeStatus) error
	GetCodeStatusByID(code authRepo.AuthorizationCode) (authRepo.CodeStatus, error)
	IsCodeExists(id int) (bool, error)
}

type Authorization interface {
	CreateUser(u authRepo.User) (authRepo.AuthorizationCode, error)
}

type AuthCodeService struct {
	repo authRepo.AuthCodeRepo
}

type AuthorizationService struct {
	repo authRepo.AuthorizationRepo
}
