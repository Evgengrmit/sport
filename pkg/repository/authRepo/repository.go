package authRepo

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"regexp"
	"time"
)

func NewAuthorization(db *sqlx.DB) *Authorization {
	return &Authorization{db: db}
}

func (a *Authorization) CreateUser(u User) (AuthorizationCode, error) {
	authCode := NewAuthCode(a.db)
	if u.Email != "" && u.Phone == "" {

		statusEmail, err := a.IsUserExistsByEmail(u)
		if err != nil {
			return AuthorizationCode{}, err
		}
		if statusEmail {
			return AuthorizationCode{}, errors.New("user with this email exists")
		}
		err = a.db.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1,$2) RETURNING id",
			u.Name, u.Email).Scan(&u.ID)
		if err != nil {
			return AuthorizationCode{}, err
		}
		return authCode.CreateCode(u)
	} else if u.Email == "" && u.Phone != "" {
		matched, err := regexp.MatchString(`^(7\d{10})?$`, u.Phone)
		if !matched || err != nil {
			return AuthorizationCode{}, errors.New("incorrect phone number format")
		}
		statusPhone, err := a.IsUserExistsByPhone(u)
		if err != nil {
			return AuthorizationCode{}, err
		}
		if statusPhone {
			return AuthorizationCode{}, errors.New("user with this phone exists")
		}
		err = a.db.DB.QueryRow("INSERT INTO users (name, phone) VALUES ($1,$2) RETURNING id",
			u.Name, u.Phone).Scan(&u.ID)
		if err != nil {
			return AuthorizationCode{}, err
		}
		return authCode.CreateCode(u)
	} else if u.Email != "" && u.Phone != "" {
		matched, err := regexp.MatchString(`^(7\d{10})?$`, u.Phone)
		if !matched || err != nil {
			return AuthorizationCode{}, errors.New("incorrect phone number format")
		}
		statusPhone, err := a.IsUserExistsByPhone(u)
		if err != nil {
			return AuthorizationCode{}, err
		}
		if statusPhone {
			return AuthorizationCode{}, errors.New("user with this phone exists")
		}
		statusEmail, err := a.IsUserExistsByEmail(u)
		if err != nil {
			return AuthorizationCode{}, err
		}
		if statusEmail {
			return AuthorizationCode{}, errors.New("user with this email exists")
		}
		err = a.db.DB.QueryRow("INSERT INTO users (name, email,phone) VALUES ($1,$2,$3) RETURNING id",
			u.Name, u.Email, u.Phone).Scan(&u.ID)
		if err != nil {
			return AuthorizationCode{}, err
		}
		return authCode.CreateCode(u)
	}
	return AuthorizationCode{}, errors.New("there must be a phone number or email")
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

func NewAuthCode(db *sqlx.DB) *AuthCode {
	return &AuthCode{db: db}
}
func (a *AuthCode) CreateCode(user User) (AuthorizationCode, error) {
	var authType string
	if user.Email != "" {
		authType = "email"
	} else if user.Phone != "" {
		authType = "phone"
	}
	rand.Seed(time.Now().UnixNano())
	authCode := AuthorizationCode{}
	authCode.Code = 1000 + rand.Intn(9000) // 1000 ≤ n ≤ 9999
	err := a.db.DB.QueryRow("INSERT INTO authorization_code (user_id, code, auth_type) VALUES ($1,$2,$3) RETURNING  id", user.ID, authCode.Code, authType).Scan(&authCode.ID)
	if err != nil {
		return AuthorizationCode{}, err
	}
	return authCode, nil
}

func (a *AuthCode) VerifyCode(code AuthorizationCode) (User, error) {
	return User{}, nil
}
