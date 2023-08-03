package user

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
)

type Service interface {
	List() ([]entity.User, error)
	Find(id string) (*entity.User, error)
	Create(input *dto.UserCreateDto) (*entity.User, int, error)
	Update(id string, input *dto.UserUpdateDto) (*entity.User, int, error)
	Delete(id string) error
}

type service struct {
	repository repository.UserRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewUserService(
	repository repository.UserRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) List() ([]entity.User, error) {
	return s.repository.List()
}

func (s *service) Find(id string) (*entity.User, error) {
	return s.repository.Find(id)
}

func (s *service) Create(input *dto.UserCreateDto) (*entity.User, int, error) {
	user := entity.User{
		ID:    s.identifier.NewUuid(),
		Name:  input.Name,
		Email: input.Email,
	}

	if err := s.validator.Validate(user); err != nil {
		return &user, customstatus.ErrBadRequest.Code, err
	}

	err := s.repository.Create(&user)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return &user, customstatus.StatusCreated.Code, nil
}

func (s *service) Update(id string, input *dto.UserUpdateDto) (*entity.User, int, error) {
	user, err := s.repository.Find(id)
	if err != nil {
		return nil, customstatus.ErrNotFound.Code, errors.New(customstatus.ErrNotFound.Message)
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := s.repository.Update(user); err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return user, customstatus.StatusCreated.Code, nil
}

func (s *service) Delete(id string) error {
	user, err := s.repository.Find(id)
	if err != nil {
		return errors.New(customstatus.ErrNotFound.Message)
	}

	return s.repository.Delete(user)
}
