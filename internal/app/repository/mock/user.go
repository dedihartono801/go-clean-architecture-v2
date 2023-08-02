package mock

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/stretchr/testify/mock"
)

// Define the MockAdminRepository interface
type MockUserRepository interface {
	Find(id string) (entity.User, error)
	Create(admin *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() MockUserRepository {
	return &mockUserRepository{}

}

// List is the mock implementation of UserRepository List method
func (m *mockUserRepository) List() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

// Find is the mock implementation of UserRepository Find method
func (m *mockUserRepository) Find(id string) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

// Create is the mock implementation of UserRepository Create method
func (m *mockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Update is the mock implementation of UserRepository Update method
func (m *mockUserRepository) Update(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Delete is the mock implementation of UserRepository Delete method
func (m *mockUserRepository) Delete(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}
