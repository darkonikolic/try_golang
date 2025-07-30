package service

import (
	"github.com/darkonikolic/try_golang/internal/domain/user/entity"
	"github.com/darkonikolic/try_golang/internal/domain/user/repository"
	"testing"
)

// MockUserRepository for testing
type MockUserRepository struct {
	users map[entity.UserID]*entity.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[entity.UserID]*entity.User),
	}
}

func (m *MockUserRepository) Save(user *entity.User) error {
	if user == nil {
		return repository.ErrInvalidUser
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) FindByID(id entity.UserID) (*entity.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, repository.ErrUserNotFound
	}
	return user, nil
}

func (m *MockUserRepository) FindByEmail(email entity.Email) (*entity.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, repository.ErrUserNotFound
}

func (m *MockUserRepository) Update(user *entity.User) error {
	if user == nil {
		return repository.ErrInvalidUser
	}

	_, exists := m.users[user.ID]
	if !exists {
		return repository.ErrUserNotFound
	}

	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) Delete(id entity.UserID) error {
	_, exists := m.users[id]
	if !exists {
		return repository.ErrUserNotFound
	}

	delete(m.users, id)
	return nil
}

func TestUserService_CreateUser(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	user, err := service.CreateUser("test@example.com", "Test User")
	if err != nil {
		t.Errorf("CreateUser() unexpected error: %v", err)
	}

	if user.Email != "test@example.com" {
		t.Errorf("CreateUser() email mismatch, got: %s", user.Email)
	}

	if user.Name != "Test User" {
		t.Errorf("CreateUser() name mismatch, got: %s", user.Name)
	}
}

func TestUserService_CreateUserWithInvalidEmail(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	_, err := service.CreateUser("invalid-email", "Test User")
	if err == nil {
		t.Errorf("CreateUser() expected error for invalid email")
	}
}

func TestUserService_CreateUserWithEmptyName(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	_, err := service.CreateUser("test@example.com", "")
	if err == nil {
		t.Errorf("CreateUser() expected error for empty name")
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	// Create user first
	user, _ := service.CreateUser("test@example.com", "Test User")

	// Get user by ID
	foundUser, err := service.GetUserByID(user.ID)
	if err != nil {
		t.Errorf("GetUserByID() unexpected error: %v", err)
	}

	if foundUser.ID != user.ID {
		t.Errorf("GetUserByID() user ID mismatch")
	}
}

func TestUserService_GetUserByIDNotFound(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	_, err := service.GetUserByID("non-existent-id")
	if err == nil {
		t.Errorf("GetUserByID() expected error, got nil")
	}
	// Check that error contains the original error
	if err.Error() != "failed to get user by ID: user not found" {
		t.Errorf("GetUserByID() unexpected error message: %v", err)
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	// Create user first
	user, _ := service.CreateUser("test@example.com", "Test User")

	// Get user by email
	foundUser, err := service.GetUserByEmail("test@example.com")
	if err != nil {
		t.Errorf("GetUserByEmail() unexpected error: %v", err)
	}

	if foundUser.Email != user.Email {
		t.Errorf("GetUserByEmail() user email mismatch")
	}
}

func TestUserService_GetUserByEmailNotFound(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	_, err := service.GetUserByEmail("notfound@example.com")
	if err == nil {
		t.Errorf("GetUserByEmail() expected error, got nil")
	}
	// Check that error contains the original error
	if err.Error() != "failed to get user by email: user not found" {
		t.Errorf("GetUserByEmail() unexpected error message: %v", err)
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	// Create user first
	user, _ := service.CreateUser("test@example.com", "Test User")

	// Update user
	err := service.UpdateUser(user.ID, "updated@example.com", "Updated Name")
	if err != nil {
		t.Errorf("UpdateUser() unexpected error: %v", err)
	}

	// Verify update
	updatedUser, _ := service.GetUserByID(user.ID)
	if updatedUser.Email != "updated@example.com" {
		t.Errorf("UpdateUser() email not updated, got: %s", updatedUser.Email)
	}

	if updatedUser.Name != "Updated Name" {
		t.Errorf("UpdateUser() name not updated, got: %s", updatedUser.Name)
	}
}

func TestUserService_UpdateUserNotFound(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	err := service.UpdateUser("non-existent-id", "updated@example.com", "Updated Name")
	if err == nil {
		t.Errorf("UpdateUser() expected error, got nil")
	}
	// Check that error contains the original error
	if err.Error() != "failed to find user for update: user not found" {
		t.Errorf("UpdateUser() unexpected error message: %v", err)
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	// Create user first
	user, _ := service.CreateUser("test@example.com", "Test User")

	// Delete user
	err := service.DeleteUser(user.ID)
	if err != nil {
		t.Errorf("DeleteUser() unexpected error: %v", err)
	}

	// Verify deletion
	_, err = service.GetUserByID(user.ID)
	if err == nil {
		t.Errorf("DeleteUser() expected error after deletion, got nil")
	}
	if err != nil && err.Error() != "failed to get user by ID: user not found" {
		t.Errorf("DeleteUser() unexpected error message after deletion: %v", err)
	}
}

func TestUserService_DeleteUserNotFound(t *testing.T) {
	repo := NewMockUserRepository()
	service := NewUserService(repo)

	err := service.DeleteUser("non-existent-id")
	if err == nil {
		t.Errorf("DeleteUser() expected error, got nil")
	}
	// Check that error contains the original error
	if err.Error() != "failed to find user for deletion: user not found" {
		t.Errorf("DeleteUser() unexpected error message: %v", err)
	}
}
