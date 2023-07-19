package graphql

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/graphql/user"
	"github.com/graphql-go/graphql"
)

type UserResovler interface {
	ResolveUsers(p graphql.ResolveParams) (interface{}, error)
	ResolveUser(p graphql.ResolveParams) (interface{}, error)
}

type userResolver struct {
	service user.Service
}

func NewUserResolver(service user.Service) UserResovler {
	return &userResolver{service}
}

func (h *userResolver) ResolveUsers(p graphql.ResolveParams) (interface{}, error) {
	users, err := h.service.GetUsers(p)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (h *userResolver) ResolveUser(p graphql.ResolveParams) (interface{}, error) {
	user, err := h.service.GetUser(p)
	if err != nil {
		return nil, err
	}
	return user, nil
}
