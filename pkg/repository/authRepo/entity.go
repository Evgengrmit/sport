package authRepo

import "github.com/jmoiron/sqlx"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}
type AuthorizationRepo interface {
	CreateUser(u User) (User, error)
	IsUserExistsByEmail(u User) (bool, error)
	IsUserExistsByID(uID int) (bool, error)
	IsUserExistsByPhone(u User) (bool, error)
}

type Authorization struct {
	db *sqlx.DB
}
