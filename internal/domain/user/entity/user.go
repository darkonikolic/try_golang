package entity

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// User represents a user entity
type User struct {
	ID        UserID
	Email     Email
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserID represents a user identifier
type UserID string

// Email represents a user email address
type Email string

// Common errors
var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrEmptyName    = errors.New("name cannot be empty")
)

// NewUser creates a new user with validation
func NewUser(email string, name string) (*User, error) {
	// Validate email
	emailObj := Email(email)
	if err := emailObj.Validate(); err != nil {
		return nil, fmt.Errorf("email validation failed: %w", err)
	}

	// Validate name
	if name == "" {
		return nil, ErrEmptyName
	}

	now := time.Now()
	user := &User{
		ID:        UserID(fmt.Sprintf("user_%d", now.UnixNano())), // Simple ID generation
		Email:     emailObj,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return user, nil
}

// Validate validates the user entity
func (u *User) Validate() error {
	if err := u.Email.Validate(); err != nil {
		return fmt.Errorf("user email validation failed: %w", err)
	}

	if u.Name == "" {
		return ErrEmptyName
	}

	return nil
}

// Update updates user information
func (u *User) Update(email string, name string) error {
	// Validate new email
	emailObj := Email(email)
	if err := emailObj.Validate(); err != nil {
		return fmt.Errorf("email validation failed: %w", err)
	}

	// Validate new name
	if name == "" {
		return ErrEmptyName
	}

	u.Email = emailObj
	u.Name = name
	u.UpdatedAt = time.Now()

	return nil
}

// IsActive checks if the user is active
func (u *User) IsActive() bool {
	// For now, all users are considered active
	// In the future, we could add an IsActive field or check against a blacklist
	return true
}

// Validate validates email format
func (e Email) Validate() error {
	if e == "" {
		return ErrInvalidEmail
	}

	// Simple email regex pattern
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(string(e)) {
		return ErrInvalidEmail
	}

	return nil
}

// String returns the email as string
func (e Email) String() string {
	return string(e)
}

// String returns the user ID as string
func (id UserID) String() string {
	return string(id)
}
