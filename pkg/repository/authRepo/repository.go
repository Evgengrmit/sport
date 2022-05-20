package authRepo

import (
	"database/sql"
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
func (a *Authorization) GetUserById(id int) (User, error) {
	exists, err := a.IsUserExistsByID(id)
	if err != nil {
		return User{}, err
	}
	if exists {
		var user User
		var phone, email sql.NullString
		err = a.db.DB.QueryRow("SELECT id, name,email,phone FROM users WHERE id =$1", id).Scan(&user.ID, &user.Name, &email, &phone)
		if err != nil {
			return User{}, err
		}
		user.Phone = phone.String
		user.Email = email.String
		return user, nil
	}
	return User{}, errors.New("there is no user with this name")
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
func (a *AuthCode) UpdateCode(code CodeStatus) error {
	if code.UsedAt.Valid {
		_, err := a.db.DB.Exec("UPDATE authorization_code SET attempt_count =$1, used_at=$2 WHERE id = $3", code.AttemptCount, code.UsedAt.Time, code.ID)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := a.db.DB.Exec("UPDATE authorization_code SET attempt_count =$1  WHERE id = $2", code.AttemptCount, code.ID)
	if err != nil {
		return err
	}

	return nil
}
func (a *AuthCode) VerifyCode(code AuthorizationCode) (User, error) {
	codeStatus, err := a.GetCodeStatusByID(code)
	if err != nil {
		return User{}, err
	}
	if codeStatus.AttemptCount >= 3 {
		return User{}, errors.New("the limit of attempts has been reached")
	}
	if codeStatus.Code != code.Code {
		codeStatus.AttemptCount += 1
		err = a.UpdateCode(codeStatus)
		if err != nil {
			return User{}, err
		}
		return User{}, errors.New("invalid code")
	}
	if codeStatus.UsedAt.Valid {
		return User{}, errors.New("this code has already been used")
	}
	if codeStatus.ExpiredAt.Before(time.Now()) {
		return User{}, errors.New("the code has expired")
	}
	codeStatus.UsedAt.Time = time.Now()
	codeStatus.UsedAt.Valid = true

	err = a.UpdateCode(codeStatus)
	if err != nil {
		return User{}, err
	}
	auth := NewAuthorization(a.db)
	user, err := auth.GetUserById(codeStatus.UserID)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (a *AuthCode) GetCodeStatusByID(code AuthorizationCode) (CodeStatus, error) {
	exists, err := a.IsCodeExists(code.ID)
	if err != nil {
		return CodeStatus{}, err
	}
	if exists {
		var codeSt CodeStatus
		err = a.db.DB.QueryRow("SELECT id, user_id, code, attempt_count, expired_at, used_at FROM authorization_code WHERE id=$1",
			code.ID).Scan(&codeSt.ID, &codeSt.UserID, &codeSt.Code, &codeSt.AttemptCount, &codeSt.ExpiredAt, &codeSt.UsedAt)
		if err != nil {
			return CodeStatus{}, err
		}
		return codeSt, nil
	}

	return CodeStatus{}, errors.New("incorrect id")
}

func (a *AuthCode) IsCodeExists(id int) (bool, error) {
	var exists bool
	err := a.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM authorization_code WHERE id = $1)",
		id).Scan(&exists)
	return exists, err
}
