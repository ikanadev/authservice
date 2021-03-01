package actions

import (
	"context"

	"github.com/vmkevv/authservice/ent"
	"github.com/vmkevv/authservice/ent/user"
)

// User all user actions
type User struct {
	ctx context.Context
	ent *ent.Client
}

// SetUpUser creates an instance of User
func SetUpUser(ctx context.Context, client *ent.Client) User {
	return User{ctx, client}
}

// Register register a new user with basic info
func (u User) Register(name, lastName, email string) (*ent.User, error) {
	return u.ent.User.Create().SetName(name).SetLastName(lastName).SetEmail(email).Save(u.ctx)
}

// ExistEmail check if an email already exists in database
func (u User) ExistEmail(email string) (bool, error) {
	return u.ent.User.Query().Where(user.EmailEQ(email)).Exist(u.ctx)
}

// GetUserByEmail get user by email
func (u User) GetUserByEmail(email string) (*ent.User, error) {
	return u.ent.User.Query().Where(user.EmailEQ(email)).First(u.ctx)
}

// GetUserByID get user info by ID
func (u User) GetUserByID(ID int) (*ent.User, error) {
	return u.ent.User.Get(u.ctx, ID)
}
