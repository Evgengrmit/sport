package authRepo

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"regexp"
)

func NewAuthorization(db *sqlx.DB) *Authorization {
	return &Authorization{db: db}
}

func (a *Authorization) CreateUser(u User) (User, error) {
	if u.Email != "" && u.Phone == "" {

		statusEmail, err := a.IsUserExistsByEmail(u)
		if err != nil {
			return User{}, err
		}
		if statusEmail {
			return User{}, errors.New("user with this email exists")
		}
		err = a.db.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1,$2) RETURNING id",
			u.Name, u.Email).Scan(&u.ID)
		return u, err
	} else if u.Email == "" && u.Phone != "" {
		matched, err := regexp.MatchString(`^(7\d{10})?$`, u.Phone)
		if !matched || err != nil {
			return User{}, errors.New("incorrect phone number format")
		}
		statusPhone, err := a.IsUserExistsByPhone(u)
		if err != nil {
			return User{}, err
		}
		if statusPhone {
			return User{}, errors.New("user with this phone exists")
		}
		err = a.db.DB.QueryRow("INSERT INTO users (name, phone) VALUES ($1,$2) RETURNING id",
			u.Name, u.Phone).Scan(&u.ID)
		return u, err
	} else if u.Email != "" && u.Phone != "" {
		matched, err := regexp.MatchString(`^(7\d{10})?$`, u.Phone)
		if !matched || err != nil {
			return User{}, errors.New("incorrect phone number format")
		}
		statusPhone, err := a.IsUserExistsByPhone(u)
		if err != nil {
			return User{}, err
		}
		if statusPhone {
			return User{}, errors.New("user with this phone exists")
		}
		statusEmail, err := a.IsUserExistsByEmail(u)
		if err != nil {
			return User{}, err
		}
		if statusEmail {
			return User{}, errors.New("user with this email exists")
		}
		err = a.db.DB.QueryRow("INSERT INTO users (name, email,phone) VALUES ($1,$2,$3) RETURNING id",
			u.Name, u.Email, u.Phone).Scan(&u.ID)
		return u, err
	}
	return User{}, errors.New("there must be a phone number or email")
}

func (a *Authorization) IsUserExistsByEmail(u User) (bool, error) {
	var exists bool
	err := a.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE email = $1)",
		u.Email).Scan(&exists)
	return exists, err
}

func (a *Authorization) IsUserExistsByPhone(u User) (bool, error) {
	var exists bool
	err := a.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE phone = $1)",
		u.Phone).Scan(&exists)
	return exists, err
}

func (a *Authorization) IsUserExistsByID(uID int) (bool, error) {
	var exists bool
	err := a.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE id = $1)",
		uID).Scan(&exists)
	return exists, err
}
