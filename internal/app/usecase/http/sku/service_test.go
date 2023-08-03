package sku

import (
	"errors"
	"testing"

	repoMock "github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository/mock"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	repo := repoMock.NewMockSkuRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewSkuService(repo, validator, identifier)

	request := &entity.Sku{
		ID:    "1",
		Name:  "kopi",
		Stock: 1,
		Price: 1,
	}

	var expected []entity.Sku
	expected = append(expected, *request)

	// Test retrieving an existing admin
	actual, err := srv.List()
	assert.Equal(t, expected[0].ID, actual[0].ID, "Expected and actual data should be equal")
	assert.Equal(t, expected[0].Name, actual[0].Name, "Expected and actual data should be equal")
	assert.Equal(t, expected[0].Stock, actual[0].Stock, "Expected and actual data should be equal")
	assert.Equal(t, expected[0].Price, actual[0].Price, "Expected and actual data should be equal")
	assert.Equal(t, nil, err, "Expected and actual data should be equal")

}

func TestCreate(t *testing.T) {
	repo := repoMock.NewMockSkuRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewSkuService(repo, validator, identifier)

	expected := &dto.SkuCreateDto{
		Name:  "Kopi",
		Stock: 1,
		Price: 10000,
	}

	expectedFailed := &dto.SkuCreateDto{
		Name:  "Failed product",
		Stock: 1,
		Price: 10000,
	}

	// Define test cases
	testCases := []struct {
		name       string
		input      *dto.SkuCreateDto
		statusCode int
		wantErr    error
	}{
		{
			name:       "Success Create data",
			input:      expected,
			statusCode: 201,
			wantErr:    nil,
		},
		{
			name:       "Failed Create data",
			input:      expectedFailed,
			statusCode: 500,
			wantErr:    errors.New("Internal Server Error"),
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call function
			rest, statusCode, err := srv.Create(tc.input)
			//fmt.Println(err.Error())
			assert.Equal(t, tc.statusCode, statusCode, "Expected status code and actual status code should be equal")
			assert.Equal(t, tc.wantErr, err, "Expected error and actual error should be equal")
			if tc.name == "Success Create data" {
				assert.Equal(t, expected.Name, rest.Name, "Expected name and actual name should be equal")
				assert.Equal(t, expected.Stock, rest.Stock, "Expected email and actual email should be equal")
				assert.Equal(t, expected.Price, rest.Price, "Expected email and actual email should be equal")
			}
		})
	}

}
