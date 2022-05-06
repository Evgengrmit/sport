package authRepo

import "github.com/jmoiron/sqlx"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type AuthorizationRepo interface {
	CreateUser(u User) (User, error)
	IsUserExists(u User) (bool, error)
}

type Authorization struct {
	db *sqlx.DB
}
