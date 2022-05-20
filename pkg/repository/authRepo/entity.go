package authRepo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
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
type CodeStatus struct {
	AuthorizationCode
	UserID       int
	AttemptCount int
	ExpiredAt    time.Time
	UsedAt       sql.NullTime
}

type AuthorizationRepo interface {
	CreateUser(u User) (AuthorizationCode, error)
	IsUserExistsByEmail(u User) (bool, error)
	IsUserExistsByID(uID int) (bool, error)
	IsUserExistsByPhone(u User) (bool, error)
	GetUserById(id int) (User, error)
}

type AuthCodeRepo interface {
	CreateCode(user User) (AuthorizationCode, error)
	UpdateCode(code CodeStatus) error
	VerifyCode(code AuthorizationCode) (User, error)
	GetCodeStatusByID(code AuthorizationCode) (CodeStatus, error)
	IsCodeExists(id int) (bool, error)
}

type Authorization struct {
	db *sqlx.DB
}

type AuthCode struct {
	db *sqlx.DB
}
