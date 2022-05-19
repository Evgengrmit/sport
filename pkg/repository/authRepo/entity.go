package authRepo

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type AuthorizationCode struct {
	ID   int `json:"authSessionId"`
	Code int `json:"code"`
}

type AuthorizationRepo interface {
	CreateUser(u User) (AuthorizationCode, error)
	IsUserExistsByEmail(u User) (bool, error)
	IsUserExistsByID(uID int) (bool, error)
	IsUserExistsByPhone(u User) (bool, error)
}

type AuthCodeRepo interface {
	CreateCode(user User) (AuthorizationCode, error)
	VerifyCode(code AuthorizationCode) (User, error)
}

type Authorization struct {
	db *sqlx.DB
}

type AuthCode struct {
	db *sqlx.DB
}
