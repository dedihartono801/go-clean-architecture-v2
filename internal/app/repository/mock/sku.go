package mock

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"gorm.io/gorm"
)

// Define the MockAdminRepository interface
type MockSkuRepository interface {
	Create(sku *entity.Sku) error
	List() ([]entity.Sku, error)
	GetSkuById(tx *gorm.DB, id string) (entity.Sku, error)
	UpdateStockSku(tx *gorm.DB, sku *entity.Sku) error
}

type mockSkuRepository struct{}

func NewMockSkuRepository() MockSkuRepository {
	return &mockSkuRepository{}

}

// List returns a list of all skus in the repository
func (m *mockSkuRepository) List() ([]entity.Sku, error) {
	var sku []entity.Sku

	dt := entity.Sku{
		ID:    "1",
		Name:  "kopi",
		Stock: 1,
		Price: 1,
	}
	sku = append(sku, dt)
	return sku, nil
}

// Create adds a new sku to the repository
func (m *mockSkuRepository) Create(sku *entity.Sku) error {
	if sku.Name == "Failed product" {
		return errors.New("failed create sku")
	}
	return nil
}

func (m *mockSkuRepository) GetSkuById(tx *gorm.DB, id string) (entity.Sku, error) {
	return entity.Sku{}, nil
}

func (m *mockSkuRepository) UpdateStockSku(tx *gorm.DB, sku *entity.Sku) error {
	return nil
}
