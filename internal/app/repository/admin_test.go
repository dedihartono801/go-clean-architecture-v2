package repository

import (
	"testing"

	//"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MockDb(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mockQ, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	// Set up GORM with the mock database
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Error opening GORM DB: %v", err)
	}
	return gdb, mockQ, err

}

func TestFind(t *testing.T) {
	// db mock
	db, mockQ, err := MockDb(t)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Create a mock admin record
	expected := &entity.Admin{
		ID:       "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(expected.ID, expected.Name, expected.Email, expected.Password)
	mockQ.ExpectQuery("SELECT * FROM `admins` WHERE `admins`.`id` = ? ORDER BY `admins`.`id` LIMIT 1").WithArgs(expected.ID).WillReturnRows(rows)

	repo := NewAdminRepository(db)

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
			// Call function
			actual, err := repo.Find(tc.id)
			if tc.name == "Found data" {
				assert.Equal(t, tc.expected, actual, "Expected and actual data should be equal")
			}

			assert.Equal(t, tc.wantErr, err != nil, "Expected error and actual error should be equal")
		})
	}
}

func TestCreate(t *testing.T) {
	// db mock
	db, mock, err := MockDb(t)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewAdminRepository(db)

	// Create a mock admin record
	now := time.Now()
	expected := &entity.Admin{
		ID:        "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:      "diding",
		Email:     "diding@gmail.com",
		Password:  "rtfgcv@098",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mock.ExpectBegin()
	// mock the INSERT query
	mock.ExpectExec("INSERT INTO `admins` (`id`,`name`,`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)").WithArgs(expected.ID, expected.Name, expected.Email, expected.Password, expected.CreatedAt, expected.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Define test cases
	testCases := []struct {
		name     string
		expected *entity.Admin
		wantErr  bool
	}{
		{
			name:     "Create data",
			expected: expected,
			wantErr:  false,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call function
			err := repo.Create(tc.expected)
			//fmt.Println(err.Error())
			assert.Equal(t, tc.wantErr, err != nil, "Expected error and actual error should be equal")
		})
	}
}
