package book

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
)

type Service interface {
	List() ([]*entity.Book, error)
	Find(id string) (*entity.Book, error)
	Create(input *dto.BookCreateDto) (*entity.Book, error)
	Update(id string, input *dto.BookUpdateDto) (*entity.Book, error)
	Delete(id string) error
}

type service struct {
	repository repository.BookRepository
	validator  validator.Validator
}

func NewService(repository repository.BookRepository, validator validator.Validator) Service {
	return &service{repository: repository, validator: validator}
}

func (s service) List() ([]*entity.Book, error) {
	return s.repository.List()
}

func (s service) Find(id string) (*entity.Book, error) {
	return s.repository.Find(id)
}

func (s service) Create(input *dto.BookCreateDto) (*entity.Book, error) {
	book := entity.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	if err := s.validator.Validate(book); err != nil {
		return &book, err
	}

	err := s.repository.Create(&book)
	return &book, err
}

func (s service) Update(id string, input *dto.BookUpdateDto) (*entity.Book, error) {
	book, err := s.repository.Find(id)
	if err != nil {
		return nil, err
	}

	book.Title = input.Title
	book.Author = input.Author

	if err := s.repository.Update(book); err != nil {
		return nil, err
	}
	return book, err
}

func (s service) Delete(id string) error {
	book, err := s.repository.Find(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(book)
}
