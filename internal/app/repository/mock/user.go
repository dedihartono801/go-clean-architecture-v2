package mock

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
)

// Define the MockAdminRepository interface
type MockUserRepository interface {
	List() ([]entity.User, error)
	Find(id string) (*entity.User, error)
	Create(admin *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

type mockUserRepository struct {
	users map[string]*entity.User
}

func NewMockUserRepository() MockUserRepository {
	return &mockUserRepository{
		users: make(map[string]*entity.User),
	}

}

// List returns a list of all users in the repository
func (m *mockUserRepository) List() ([]entity.User, error) {
	var userList []entity.User
	for _, user := range m.users {
		userList = append(userList, *user)
	}
	return userList, nil
}

// Find retrieves a user by ID from the repository
func (m *mockUserRepository) Find(id string) (*entity.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Create adds a new user to the repository
func (m *mockUserRepository) Create(user *entity.User) error {
	if _, ok := m.users[user.ID]; ok {
		return errors.New("user already exists")
	}
	m.users[user.ID] = user
	return nil
}

// Update updates an existing user in the repository
func (m *mockUserRepository) Update(user *entity.User) error {
	if _, ok := m.users[user.ID]; !ok {
		return errors.New("user not found")
	}
	m.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository
func (m *mockUserRepository) Delete(user *entity.User) error {
	if _, ok := m.users[user.ID]; !ok {
		return errors.New("user not found")
	}
	delete(m.users, user.ID)
	return nil
}
