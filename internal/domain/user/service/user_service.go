package service

import (
	"fmt"
	"github.com/darkonikolic/try_golang/internal/domain/user/entity"
	"github.com/darkonikolic/try_golang/internal/domain/user/repository"
)

// UserService handles business logic for user operations
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser creates a new user with validation
func (s *UserService) CreateUser(email string, name string) (*entity.User, error) {
	// Check if user already exists with this email
	existingUser, err := s.repo.FindByEmail(entity.Email(email))
	if err == nil && existingUser != nil {
		return nil, repository.ErrUserAlreadyExists
	}

	// Create new user
	user, err := entity.NewUser(email, name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Save user
	if err := s.repo.Save(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id entity.UserID) (*entity.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email
func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.repo.FindByEmail(entity.Email(email))
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id entity.UserID, email string, name string) error {
	// Get existing user
	user, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find user for update: %w", err)
	}

	// Update user fields
	if err := user.Update(email, name); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	// Save updated user
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("failed to save updated user: %w", err)
	}

	return nil
}

// DeleteUser removes a user by their ID
func (s *UserService) DeleteUser(id entity.UserID) error {
	// First check if user exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find user for deletion: %w", err)
	}

	// Delete user
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// IsUserActive checks if a user is active
func (s *UserService) IsUserActive(id entity.UserID) (bool, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return false, fmt.Errorf("failed to find user: %w", err)
	}

	return user.IsActive(), nil
}
