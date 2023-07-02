package product

import (
	"encoding/json"
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
)

type Service interface {
	Product() (*dto.ProductDto, int, error)
}
type service struct {
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewProductService(
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) Product() (*dto.ProductDto, int, error) {
	prd, err := helpers.GetHTTPRequest("https://run.mocky.io/v3/70437170-6f70-4c3b-8c27-f476ff697236", "")
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	detailPrd := &dto.ProductDto{}
	err = json.Unmarshal(prd, &detailPrd)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, err
	}

	return detailPrd, customstatus.StatusOk.Code, nil
}
