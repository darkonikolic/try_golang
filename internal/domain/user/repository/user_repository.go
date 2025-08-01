package repository

import (
	"errors"
	"github.com/darkonikolic/try_golang/internal/domain/user/entity"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	// Save creates a new user or updates existing one
	Save(user *entity.User) error

	// FindByID retrieves a user by their ID
	FindByID(id entity.UserID) (*entity.User, error)

	// FindByEmail retrieves a user by their email
	FindByEmail(email entity.Email) (*entity.User, error)

	// Update updates an existing user
	Update(user *entity.User) error

	// Delete removes a user by their ID
	Delete(id entity.UserID) error
}

// Domain-specific errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUser       = errors.New("invalid user data")
)
