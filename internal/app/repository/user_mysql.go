package repository

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	List() ([]entity.User, error)
	Find(id string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{database}
}

func (r *userRepository) List() ([]entity.User, error) {
	users := []entity.User{}
	err := r.database.Find(&users).Error
	return users, err
}

func (r *userRepository) Find(id string) (*entity.User, error) {
	user := entity.User{ID: id}
	err := r.database.First(&user).Error
	return &user, err
}

func (r *userRepository) Create(user *entity.User) error {
	return r.database.Create(user).Error
}

func (r *userRepository) Update(user *entity.User) error {

	return r.database.Save(user).Error
}

func (r *userRepository) Delete(user *entity.User) error {
	return r.database.Delete(user).Error
}
