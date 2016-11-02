package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	log "github.com/iamthemuffinman/logsip"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user of our application
type User struct {
	ID         int64  `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"password,omitempty" db:"password"`
	Email      string `json:"email" db:"email"`
	FullName   string `json:"full_name" db:"full_name"`
	Gravatar   string `json:"gravatar" db:"gravatar"`
	ProfilePic string `json:"profile_pic" db:"profile_pic"`
	IsAdmin    bool   `json:"is_admin,omitempty" db:"is_admin"`
}

// CheckPw will verify if the given password matches for this user. Logs any
// errors it encounters
func (u *User) CheckPw(pw []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), pw)
	if err == nil {
		return true
	}

	log.Error(err)
	return false
}

func (u *User) String() string {
	return fmt.Sprintf("<User %d, %s>", u.ID, u.Username)
}

// NewUser will create the user after encrypting the password with bcrypt
func NewUser(username, password, fullName, email string, admin bool) (*User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	emailHash := md5.Sum([]byte(strings.ToLower(email)))
	eh := hex.EncodeToString(emailHash[:16])

	return &User{
		Username:   username,
		Password:   string(pw),
		Email:      email,
		FullName:   fullName,
		ProfilePic: "https://www.gravatar.com/avatar/" + eh,
		Gravatar:   "https://www.gravatar.com/avatar/" + eh,
		IsAdmin:    admin,
	}, nil
}
