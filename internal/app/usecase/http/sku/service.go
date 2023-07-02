package sku

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
	Create(input *dto.SkuCreateDto) (*entity.Sku, int, error)
	List() ([]entity.Sku, error)
}

type service struct {
	repository repository.SkuRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewSkuService(
	repository repository.SkuRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) Create(input *dto.SkuCreateDto) (*entity.Sku, int, error) {
	sku := entity.Sku{
		ID:    s.identifier.NewUuid(),
		Name:  input.Name,
		Stock: input.Stock,
		Price: input.Price,
	}

	if err := s.validator.Validate(sku); err != nil {
		return &sku, customstatus.ErrBadRequest.Code, err
	}

	err := s.repository.Create(&sku)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return &sku, customstatus.StatusCreated.Code, nil
}

func (s *service) List() ([]entity.Sku, error) {
	return s.repository.List()
}
