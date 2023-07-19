// memory is a in memory data storage solution for Users
package repository

import (
	"errors"
	"sync"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
)

type MemoryUserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUser(id string) (entity.User, error)
}

// InMemoryRepository is a storage for users that uses a map to store them
type memoryUserRepository struct {
	// users is our super storage for users.
	user []entity.User
	sync.Mutex
}

// NewMemoryRepository initializes a memory with mock data
func NewMemoryUserRepository() MemoryUserRepository {
	user := []entity.User{
		{
			ID:    "1",
			Name:  "Sandi",
			Email: "sandi@gmail.com",
		}, {
			ID:    "2",
			Name:  "deni",
			Email: "deni@gmail.com",
		},
	}

	return &memoryUserRepository{
		user: user,
	}
}

// GetUsers returns all users
func (imr *memoryUserRepository) GetUsers() ([]entity.User, error) {
	return imr.user, nil
}

// GetUsers will return a goper by its ID
func (imr *memoryUserRepository) GetUser(id string) (entity.User, error) {
	for _, user := range imr.user {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, errors.New("no such users exists")
}
