package repository

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SkuRepository interface {
	Create(sku *entity.Sku) error
	List() ([]entity.Sku, error)
	GetSkuById(tx *gorm.DB, id string) (entity.Sku, error)
	UpdateStockSku(tx *gorm.DB, sku *entity.Sku) error
}

type skuRepository struct {
	database *gorm.DB
}

func NewSkuRepository(database *gorm.DB) SkuRepository {
	return &skuRepository{database}
}

func (r *skuRepository) Create(sku *entity.Sku) error {
	return r.database.Table("sku").Create(sku).Error
}

func (r *skuRepository) List() ([]entity.Sku, error) {
	skus := []entity.Sku{}
	err := r.database.Table("sku").Find(&skus).Error
	return skus, err
}

func (r *skuRepository) GetSkuById(tx *gorm.DB, id string) (entity.Sku, error) {
	sku := entity.Sku{ID: id}
	err := tx.Table("sku").Clauses(clause.Locking{Strength: "UPDATE"}).First(&sku).Error
	return sku, err
}

func (r *skuRepository) UpdateStockSku(tx *gorm.DB, sku *entity.Sku) error {

	return tx.Table("sku").Save(sku).Error
}
