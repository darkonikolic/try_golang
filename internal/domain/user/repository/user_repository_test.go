package repository

import (
	"github.com/darkonikolic/try_golang/internal/domain/user/entity"
	"testing"
)

// MockUserRepository implements UserRepository for testing
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
		return ErrInvalidUser
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) FindByID(id entity.UserID) (*entity.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (m *MockUserRepository) FindByEmail(email entity.Email) (*entity.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (m *MockUserRepository) Update(user *entity.User) error {
	if user == nil {
		return ErrInvalidUser
	}

	_, exists := m.users[user.ID]
	if !exists {
		return ErrUserNotFound
	}

	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) Delete(id entity.UserID) error {
	_, exists := m.users[id]
	if !exists {
		return ErrUserNotFound
	}

	delete(m.users, id)
	return nil
}

func TestUserRepository_Save(t *testing.T) {
	repo := NewMockUserRepository()
	user, _ := entity.NewUser("test@example.com", "Test User")

	err := repo.Save(user)
	if err != nil {
		t.Errorf("Save() unexpected error: %v", err)
	}

	// Verify user was saved
	savedUser, err := repo.FindByID(user.ID)
	if err != nil {
		t.Errorf("FindByID() failed to find saved user: %v", err)
	}

	if savedUser.Email != user.Email {
		t.Errorf("Save() user email mismatch, got: %s, want: %s", savedUser.Email, user.Email)
	}
}

func TestUserRepository_SaveNilUser(t *testing.T) {
	repo := NewMockUserRepository()

	err := repo.Save(nil)
	if err != ErrInvalidUser {
		t.Errorf("Save() expected ErrInvalidUser, got: %v", err)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	repo := NewMockUserRepository()
	user, _ := entity.NewUser("test@example.com", "Test User")
	err := repo.Save(user)
	if err != nil {
		t.Errorf("Save() failed to save user: %v", err)
	}

	foundUser, err := repo.FindByID(user.ID)
	if err != nil {
		t.Errorf("FindByID() unexpected error: %v", err)
	}

	if foundUser.ID != user.ID {
		t.Errorf("FindByID() user ID mismatch")
	}
}

func TestUserRepository_FindByIDNotFound(t *testing.T) {
	repo := NewMockUserRepository()

	_, err := repo.FindByID("non-existent-id")
	if err == nil {
		t.Errorf("FindByID() expected error, got nil")
	}
	if err != ErrUserNotFound {
		t.Errorf("FindByID() expected ErrUserNotFound, got: %v", err)
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	repo := NewMockUserRepository()
	user, _ := entity.NewUser("test@example.com", "Test User")
	err := repo.Save(user)
	if err != nil {
		t.Errorf("Save() failed to save user: %v", err)
	}

	foundUser, err := repo.FindByEmail(user.Email)
	if err != nil {
		t.Errorf("FindByEmail() unexpected error: %v", err)
	}

	if foundUser.Email != user.Email {
		t.Errorf("FindByEmail() user email mismatch")
	}
}

func TestUserRepository_FindByEmailNotFound(t *testing.T) {
	repo := NewMockUserRepository()

	_, err := repo.FindByEmail("notfound@example.com")
	if err == nil {
		t.Errorf("FindByEmail() expected error, got nil")
	}
	if err != ErrUserNotFound {
		t.Errorf("FindByEmail() expected ErrUserNotFound, got: %v", err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	repo := NewMockUserRepository()
	user, _ := entity.NewUser("test@example.com", "Test User")
	err := repo.Save(user)
	if err != nil {
		t.Errorf("Save() failed to save user: %v", err)
	}

	// Update user
	err = user.Update("updated@example.com", "Updated Name")
	if err != nil {
		t.Errorf("Update() failed to update user: %v", err)
	}

	err = repo.Update(user)
	if err != nil {
		t.Errorf("Update() unexpected error: %v", err)
	}

	// Verify update
	updatedUser, _ := repo.FindByID(user.ID)
	if updatedUser.Email != "updated@example.com" {
		t.Errorf("Update() email not updated, got: %s", updatedUser.Email)
	}
}

func TestUserRepository_UpdateNotFound(t *testing.T) {
	repo := NewMockUserRepository()
	user, _ := entity.NewUser("test@example.com", "Test User")

	err := repo.Update(user)
	if err != ErrUserNotFound {
		t.Errorf("Update() expected ErrUserNotFound, got: %v", err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	repo := NewMockUserRepository()
	user, _ := entity.NewUser("test@example.com", "Test User")
	err := repo.Save(user)
	if err != nil {
		t.Errorf("Save() failed to save user: %v", err)
	}

	err = repo.Delete(user.ID)
	if err != nil {
		t.Errorf("Delete() unexpected error: %v", err)
	}

	// Verify deletion
	_, err = repo.FindByID(user.ID)
	if err != ErrUserNotFound {
		t.Errorf("Delete() user still exists after deletion")
	}
}

func TestUserRepository_DeleteNotFound(t *testing.T) {
	repo := NewMockUserRepository()

	err := repo.Delete("non-existent-id")
	if err != ErrUserNotFound {
		t.Errorf("Delete() expected ErrUserNotFound, got: %v", err)
	}
}
