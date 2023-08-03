package admin

import (
	"testing"

	repoMock "github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository/mock"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	repo := repoMock.NewMockAdminRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewAdminService(repo, validator, identifier)

	expected := &entity.Admin{
		ID:       "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
	}
	repo.Create(expected)

	// Define test cases
	testCases := []struct {
		name     string
		id       string
		expected *entity.Admin
		wantErr  bool
	}{
		{
			name:     "Found data",
			id:       expected.ID,
			expected: expected,
			wantErr:  false,
		},
		{
			name:     "Not found data",
			id:       "8734JJHYD88",
			expected: nil,
			wantErr:  true,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Test retrieving an existing admin
			actual, err := srv.Find(tc.id)
			assert.Equal(t, tc.expected, actual, "Expected and actual data should be equal")

			assert.Equal(t, tc.wantErr, err != nil, "Expected error and actual error should be equal")
		})
	}

}

func TestCreate(t *testing.T) {
	repo := repoMock.NewMockAdminRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewAdminService(repo, validator, identifier)

	expected := &dto.AdminCreateDto{
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
	}

	// Define test cases
	testCases := []struct {
		name       string
		input      *dto.AdminCreateDto
		statusCode int
		wantErr    error
	}{
		{
			name:       "Success Create data",
			input:      expected,
			statusCode: 201,
			wantErr:    nil,
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
			assert.Equal(t, expected.Name, rest.Name, "Expected name and actual name should be equal")
			assert.Equal(t, expected.Email, rest.Email, "Expected email and actual email should be equal")
		})
	}

}
