package mock

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
)

// Define the MockAdminRepository interface
type MockAdminRepository interface {
	Find(id string) (*entity.Admin, error)
	Create(admin *entity.Admin) error
	FindByEmail(email string) (*entity.Admin, error)
}

type mockAdminRepository struct {
	admin map[string]*entity.Admin
}

func NewMockAdminRepository() MockAdminRepository {
	return &mockAdminRepository{
		admin: make(map[string]*entity.Admin),
	}

}

// Create a new admin
func (m *mockAdminRepository) Create(admin *entity.Admin) error {
	if _, ok := m.admin[admin.ID]; ok {
		return errors.New("user already exists")
	}
	m.admin[admin.ID] = admin
	return nil
}

// // Update an existing user
// func (m *mockAdminRepository) Update(user *entity.User) error {
// 	if _, ok := m.users[user.ID]; !ok {
// 		return errors.New("user not found")
// 	}
// 	user.UpdatedAt = time.Now()
// 	m.users[user.ID] = user
// 	return nil
// }

// Delete an existing user
// func (m *mockAdminRepository) Delete(id string) error {
// 	if _, ok := m.users[id]; !ok {
// 		return errors.New("user not found")
// 	}
// 	delete(m.users, id)
// 	return nil
// }

// FindByID finds a user by ID
func (m *mockAdminRepository) Find(id string) (*entity.Admin, error) {
	if admin, ok := m.admin[id]; ok {
		return admin, nil
	}
	return nil, errors.New("admin not found")
}

// FindByEmail finds a user by email
func (m *mockAdminRepository) FindByEmail(email string) (*entity.Admin, error) {
	for _, admin := range m.admin {
		if admin.Email == email {
			return admin, nil
		}
	}
	return nil, errors.New("admin not found")
}
