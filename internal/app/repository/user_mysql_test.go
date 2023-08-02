package repository

import (
	"errors"
	"testing"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database for testing")
	}
	db.AutoMigrate(&entity.User{})
	return db
}

func TestUserRepository(t *testing.T) {
	// db mock
	db := setupTestDB()

	// Create the UserRepository using the test database
	userRepo := NewUserRepository(db)

	// Insert some test data into the database
	users := []*entity.User{
		{
			ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
			Name:  "diding",
			Email: "diding@gmail.com",
		},
		{
			ID:    "4d35bf38-8c50-4c85-8072-fd9794803a17",
			Name:  "deden",
			Email: "deden@gmail.com",
		},
		{
			ID:    "4d35bf38-8c50-4c85-8072-fd9794803a18",
			Name:  "dani",
			Email: "dani@gmail.com",
		},
	}

	t.Run("User Create", func(t *testing.T) {
		for _, user := range users {
			err := userRepo.Create(user)
			if err != nil {
				t.Fatalf("error creating user: %v", err)
			}
			// Find the user by ID
			foundUser, err := userRepo.Find(user.ID)
			if err != nil {
				t.Fatalf("error finding user: %v", err)
			}

			// Verify the user data
			if foundUser.ID != user.ID || foundUser.Name != user.Name || foundUser.Email != user.Email {
				t.Errorf("expected user data to match, got: %+v, want: %+v", foundUser, user)
			}
		}
	})

	t.Run("User List", func(t *testing.T) {
		// Call function
		actual, err := userRepo.List()
		if err != nil {
			t.Fatalf("error listing users: %v", err)
		}

		assert.Equal(t, len(users), len(actual), "Expected and actual data should be equal")

		for i, user := range users {
			if user.ID != actual[i].ID || user.Name != actual[i].Name || user.Email != actual[i].Email {
				t.Errorf("expected user data to match, got: %+v, want: %+v", actual[i], user)
			}
		}

	})

	t.Run("User Find", func(t *testing.T) {
		expected := &entity.User{
			ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
			Name:  "diding",
			Email: "diding@gmail.com",
		}
		// Define test cases
		testCases := []struct {
			name     string
			id       string
			expected *entity.User
			wantErr  error
		}{
			{
				name:     "Found data",
				id:       expected.ID,
				expected: expected,
				wantErr:  nil,
			},
			{
				name:     "Not found data",
				id:       "8734JJHYD88",
				expected: nil,
				wantErr:  errors.New("record not found"),
			},
		}

		// Run tests
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Call function
				actual, err := userRepo.Find(tc.id)
				if tc.name == "Not found data" {
					assert.Equal(t, tc.wantErr, err, "Expected and actual data should be equal")
				}
				if tc.name == "Found data" {
					// Verify the user data
					if tc.expected.ID != actual.ID || tc.expected.Name != actual.Name || tc.expected.Email != actual.Email {
						t.Errorf("expected user data to match, got: %+v, want: %+v", expected, actual)
					}
				}

			})
		}
	})

	t.Run("User Update", func(t *testing.T) {
		user := &entity.User{
			ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
			Name:  "Bambang",
			Email: "Bambang@gmail.com",
		}
		err := userRepo.Update(user)
		if err != nil {
			t.Fatalf("error updating user: %v", err)
		}

		foundUser, err := userRepo.Find(user.ID)
		if err != nil {
			t.Fatalf("error finding user: %v", err)
		}
		if foundUser.Name != "Bambang" {
			t.Errorf("expected updated age to be Bambang, got: %s", foundUser.Name)
		}
	})

	t.Run("User Delete", func(t *testing.T) {
		user := &entity.User{
			ID: "4d35bf38-8c50-4c85-8072-fd9794803a16",
		}
		// Delete the user
		err := userRepo.Delete(user)
		if err != nil {
			t.Fatalf("error deleting user: %v", err)
		}

		// Try to find the deleted user
		_, err = userRepo.Find(user.ID)
		if err == nil {
			t.Error("expected error when finding deleted user")
		}
	})
}
