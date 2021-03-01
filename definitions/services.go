package definitions

import "github.com/vmkevv/authservice/ent"

// UserActions represents all user posible actions
type UserActions interface {
	Register(name, lastName, email string) (*ent.User, error)
	ExistEmail(email string) (bool, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUserByID(ID int) (*ent.User, error)
}
