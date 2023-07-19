package user

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	"github.com/graphql-go/graphql"
)

type Service interface {
	GetUsers(p graphql.ResolveParams) ([]entity.User, error)
	GetUser(p graphql.ResolveParams) (*entity.User, error)
}

type service struct {
	repository repository.MemoryUserRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewGraphqlUserService(
	repository repository.MemoryUserRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) GetUsers(p graphql.ResolveParams) ([]entity.User, error) {
	return s.repository.GetUsers()
}

func (s *service) GetUser(p graphql.ResolveParams) (*entity.User, error) {

	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("id has to be a string")
	}
	user, err := s.repository.GetUser(id)
	if err != nil {
		return nil, errors.New("does not exist")
	}
	return &user, nil
}
