package repository

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/google/uuid"
)

type BookRepository interface {
	List() ([]*entity.Book, error)
	Find(id string) (*entity.Book, error)
	Create(book *entity.Book) error
	Update(book *entity.Book) error
	Delete(book *entity.Book) error
}

type bookRepository struct {
	data map[string]*entity.Book
}

func NewBookRepository() BookRepository {
	var mp = map[string]*entity.Book{}
	return &bookRepository{data: mp}
}

func (r *bookRepository) List() ([]*entity.Book, error) {
	var books []*entity.Book
	for _, book := range r.data {
		books = append(books, book)
	}
	return books, nil
}

func (r *bookRepository) Find(id string) (*entity.Book, error) {
	if r.data[id] == nil {
		return nil, errors.New(customstatus.ErrNotFound.Message)
	}
	return r.data[id], nil
}

func (r *bookRepository) Create(book *entity.Book) error {
	book.ID = uuid.New().String()
	r.data[book.ID] = book
	return nil
}

func (r *bookRepository) Update(book *entity.Book) error {
	_, err := r.Find(book.ID)
	if err != nil {
		return errors.New(customstatus.ErrNotFound.Message)
	}
	r.data[book.ID] = book
	return nil
}

func (r *bookRepository) Delete(book *entity.Book) error {
	r.data[book.ID] = nil
	return nil
}
