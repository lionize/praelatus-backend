package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user of our application
type User struct {
	ID              uint         `json:"-" gorm:"primary_key"`
	Username        string       `json:"username" gorm:"unique"`
	Password        string       `json:"password,omitempty"`
	Email           string       `json:"email" gorm:"unique"`
	FullName        string       `json:"fullName"`
	Memberships     []Membership `json:"-"`
	AssignedTickets []Ticket     `json:"assignedTickets,omitempty" gorm:"ForeignKey:Assignee"`
	ReportedTickets []Ticket     `json:"reportedTickets,omitempty" gorm:"ForeignKey:Reporter"`
	IsAdmin         bool         `json:"isAdmin"`
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

func (u *User) String() string {
	return fmt.Sprintf("<User %d, %s>", u.ID, u.Username)
}
