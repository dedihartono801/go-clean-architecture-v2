package user

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
	repo := repoMock.NewMockUserRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewUserService(repo, validator, identifier)

	request := &entity.User{
		ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:  "diding",
		Email: "diding@gmail.com",
	}
	repo.Create(request)

	var expected []entity.User
	expected = append(expected, *request)

	// Test retrieving an existing admin
	actual, err := srv.List()
	assert.Equal(t, expected, actual, "Expected and actual data should be equal")
	assert.Equal(t, nil, err, "Expected and actual data should be equal")

}

func TestFind(t *testing.T) {
	repo := repoMock.NewMockUserRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewUserService(repo, validator, identifier)

	expected := &entity.User{
		ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:  "diding",
		Email: "diding@gmail.com",
	}
	repo.Create(expected)

	// Define test cases
	testCases := []struct {
		name     string
		id       string
		expected *entity.User
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
	repo := repoMock.NewMockUserRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewUserService(repo, validator, identifier)

	expected := &dto.UserCreateDto{
		Name:  "diding",
		Email: "diding@gmail.com",
	}

	// Define test cases
	testCases := []struct {
		name       string
		input      *dto.UserCreateDto
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

func TestUpdate(t *testing.T) {
	repo := repoMock.NewMockUserRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewUserService(repo, validator, identifier)

	request := &entity.User{
		ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:  "diding",
		Email: "diding@gmail.com",
	}
	repo.Create(request)

	requestUpdate := &dto.UserUpdateDto{
		Name:  "deden",
		Email: "deden@gmail.com",
	}

	testCases := []struct {
		name       string
		ID         string
		statusCode int
		wantErr    error
	}{
		{
			name:       "Success Update data",
			ID:         "4d35bf38-8c50-4c85-8072-fd9794803a16",
			statusCode: 201,
			wantErr:    nil,
		},
		{
			name:       "Failed Update data",
			ID:         "4d35bf38-8c50-4c85-8072-ih787yb",
			statusCode: 404,
			wantErr:    errors.New("Data not found"),
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Test retrieving an existing admin
			actual, statusCode, err := srv.Update(tc.ID, requestUpdate)
			assert.Equal(t, tc.wantErr, err, "Expected and actual data should be equal")
			assert.Equal(t, tc.statusCode, statusCode, "Expected and actual data should be equal")
			if tc.name == "Success Update data" {
				assert.Equal(t, requestUpdate.Name, actual.Name, "Expected and actual data should be equal")
				assert.Equal(t, requestUpdate.Email, actual.Email, "Expected and actual data should be equal")
			}
		})
	}

}

func TestDelete(t *testing.T) {
	repo := repoMock.NewMockUserRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewUserService(repo, validator, identifier)

	request := &entity.User{
		ID:    "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:  "diding",
		Email: "diding@gmail.com",
	}
	repo.Create(request)

	testCases := []struct {
		name    string
		ID      string
		wantErr error
	}{
		{
			name:    "Success Delete data",
			ID:      "4d35bf38-8c50-4c85-8072-fd9794803a16",
			wantErr: nil,
		},
		{
			name:    "Failed Delete data",
			ID:      "4d35bf38-8c50-4c85-8072-ih787yb",
			wantErr: errors.New("Data not found"),
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Test retrieving an existing admin
			err := srv.Delete(tc.ID)
			assert.Equal(t, tc.wantErr, err, "Expected and actual data should be equal")
		})
	}

}
