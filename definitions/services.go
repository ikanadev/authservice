package definitions

import "github.com/vmkevv/authservice/ent"

// UserActions represents all user posible actions
type UserActions interface {
	Register(name, lastName, email string) (*ent.User, error)
	Update(name, lastName, profilePicture, email, phone, about string) (*ent.User, error)
	UpdateRole(role int8) (*ent.User, error)
	SendEmailToken(email string) error
	GenerateToken(ID int, role int8) (string, error)
	ExistEmail(email string) bool
	GetUserByToken(token string) (*ent.User, error)
}
