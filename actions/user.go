package actions

import (
	"context"

	"github.com/vmkevv/authservice/ent"
)

// User all user actions
type User struct {
	ctx context.Context
	ent *ent.Client
}

// Register register a new user with basic info
func (u *User) Register(name, lastName, email string) (*ent.User, error) {
	return nil, nil
}

// Update updates user info
func (u *User) Update(name, lastName, profilePicture, email, phone, about string) (*ent.User, error) {
	return nil, nil
}

// UpdateRole updates user current role
func (u *User) UpdateRole(role int8) (*ent.User, error) {
	return nil, nil
}

// SendEmailToken sends the login magic link to user email
func (u *User) SendEmailToken(email string) error {
	return nil
}

// GenerateToken generates a user token based in user ID and role
func (u *User) GenerateToken(ID int, role int8) (string, error) {
	return "", nil
}

// ExistEmail check if an email already exists in database
func (u *User) ExistEmail(email string) bool {
	return false
}

// GetUserByToken get user info by token
func (u *User) GetUserByToken(token string) (*ent.User, error) {
	return nil, nil
}
