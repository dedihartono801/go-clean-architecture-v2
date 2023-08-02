package http

import (
	"testing"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	repoMock "github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository/mock"
	usecaseMock "github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/admin/mock"
	adminDto "github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func MockNewService(t *testing.T) usecaseMock.Service {

	repo := repoMock.NewMockAdminRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := usecaseMock.NewMockAdminService(repo, validator, identifier)
	return srv
}

func TestFind(t *testing.T) {
	srv := MockNewService(t)

	// Create a new Fiber app
	app := fiber.New()

	// Define test cases
	testCases := []struct {
		name       string
		id         string
		statusCode int
	}{
		{
			name:       "Found data",
			id:         "4d35bf38-8c50-4c85-8072-fd9794803a167",
			statusCode: 200,
		},
		{
			name:       "Not found data",
			id:         "8734JJHYD88",
			statusCode: 404,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Define the getUserByID route
			app.Get("/admin", func(c *fiber.Ctx) error {
				// Set the adminID value in the context's locals map
				c.Locals("adminID", tc.id)
				handler := NewAdminHandler(srv)
				return handler.Find(c)
			})

			// Define a mock request for testing
			req := httptest.NewRequest(http.MethodGet, "/admin", nil)

			resp, err := app.Test(req, -1)

			// ensure that there are no errors
			assert.NoError(t, err)

			// ensure that the response status code is 200 OK
			assert.Equal(t, tc.statusCode, resp.StatusCode)
		})
	}

}

func TestCreate(t *testing.T) {
	srv := MockNewService(t)

	// Create a new Fiber app
	app := fiber.New()

	// Define test cases
	input := &adminDto.AdminCreateDto{
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "rtdfxc@123",
	}
	// Define test cases
	testCases := []struct {
		name       string
		input      *adminDto.AdminCreateDto
		statusCode int
		wantErr    bool
	}{
		{
			name:       "Success Create data",
			input:      input,
			statusCode: 201,
			wantErr:    false,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Define the getUserByID route
			reqBodyBytes, _ := json.Marshal(tc.input)

			app.Post("/admin/create", func(c *fiber.Ctx) error {
				// Set the adminID value in the context's locals map
				handler := NewAdminHandler(srv)
				return handler.Create(c)
			})

			req := httptest.NewRequest(http.MethodPost, "/admin/create", bytes.NewBuffer(reqBodyBytes))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Error testing Create User: %v", err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != tc.statusCode {
				t.Errorf("Expected status code %d but got %d", tc.statusCode, resp.StatusCode)
			}
		})
	}
}
