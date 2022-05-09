package authRepo

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

func NewAuthorization(db *sqlx.DB) *Authorization {
	return &Authorization{db: db}
}

func (a *Authorization) CreateUser(u User) (User, error) {
	status, err := a.IsUserExists(u)
	if err != nil {
		return User{}, err
	}
	if status {
		return User{}, errors.New("user with this email exists")
	}
	err = a.db.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1,$2) RETURNING id",
		u.Name, u.Email).Scan(&u.ID)
	return u, err
}

func (a *Authorization) IsUserExists(u User) (bool, error) {
	var exists bool
	err := a.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE email = $1)",
		u.Email).Scan(&exists)
	return exists, err
}

func (a *Authorization) IsUserExistsByID(uID int) (bool, error) {
	var exists bool
	err := a.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE id = $1)",
		uID).Scan(&exists)
	return exists, err
}
