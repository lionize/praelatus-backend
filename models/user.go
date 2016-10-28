package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user of our application
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	FullName string `json:"full_name" db:"full_name"`
	IsAdmin  bool   `json:"is_admin,omitempty" db:"is_admin"`
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

	return &User{
		Username: username,
		Password: string(pw),
		Email:    email,
		FullName: fullName,
		IsAdmin:  admin,
	}, nil
}
