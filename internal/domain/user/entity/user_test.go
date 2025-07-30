package entity

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		userName string
		wantErr  bool
	}{
		{
			name:     "valid user",
			email:    "test@example.com",
			userName: "Test User",
			wantErr:  false,
		},
		{
			name:     "invalid email",
			email:    "invalid-email",
			userName: "Test User",
			wantErr:  true,
		},
		{
			name:     "empty name",
			email:    "test@example.com",
			userName: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.email, tt.userName)

			if tt.wantErr && err == nil {
				t.Errorf("NewUser() expected error but got none")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("NewUser() unexpected error: %v", err)
			}

			if !tt.wantErr && user == nil {
				t.Errorf("NewUser() returned nil user when no error expected")
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	user, _ := NewUser("test@example.com", "Test User")

	if err := user.Validate(); err != nil {
		t.Errorf("User.Validate() unexpected error: %v", err)
	}
}

func TestUser_Update(t *testing.T) {
	user, _ := NewUser("test@example.com", "Test User")
	originalUpdatedAt := user.UpdatedAt

	// Wait a bit to ensure UpdatedAt will be different
	time.Sleep(1 * time.Millisecond)

	err := user.Update("new@example.com", "New Name")
	if err != nil {
		t.Errorf("User.Update() unexpected error: %v", err)
	}

	if user.Email != "new@example.com" {
		t.Errorf("User.Update() email not updated, got: %s", user.Email)
	}

	if user.Name != "New Name" {
		t.Errorf("User.Update() name not updated, got: %s", user.Name)
	}

	if user.UpdatedAt.Equal(originalUpdatedAt) {
		t.Errorf("User.Update() UpdatedAt not changed")
	}
}

func TestUser_IsActive(t *testing.T) {
	user, _ := NewUser("test@example.com", "Test User")

	if !user.IsActive() {
		t.Errorf("User.IsActive() should return true for new user")
	}
}

func TestEmail_Validate(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "test@example.com", false},
		{"valid email with subdomain", "test@sub.example.com", false},
		{"invalid email no @", "testexample.com", true},
		{"invalid email no domain", "test@", true},
		{"invalid email no local", "@example.com", true},
		{"empty email", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := Email(tt.email)
			err := email.Validate()

			if tt.wantErr && err == nil {
				t.Errorf("Email.Validate() expected error but got none")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("Email.Validate() unexpected error: %v", err)
			}
		})
	}
}
