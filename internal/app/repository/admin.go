package repository

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Find(id string) (*entity.Admin, error)
	Create(user *entity.Admin) error
	FindByEmail(email string) (*entity.Admin, error)
}

type adminRepository struct {
	database *gorm.DB
}

func NewAdminRepository(database *gorm.DB) AdminRepository {
	return &adminRepository{database}
}

func (r *adminRepository) Find(id string) (*entity.Admin, error) {
	admin := &entity.Admin{ID: id}
	err := r.database.First(&admin).Error
	return admin, err
}

func (r *adminRepository) Create(admin *entity.Admin) error {
	return r.database.Create(admin).Error
}

func (r *adminRepository) FindByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.database.Where("email = ?", email).First(&admin).Error
	return &admin, err
}
