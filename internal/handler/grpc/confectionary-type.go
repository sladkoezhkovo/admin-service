package grpc

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type ConfectionaryTypeService interface {
	Create(city *entity.ConfectionaryType) error
	FindById(id int) (*entity.ConfectionaryType, error)
	FindByName(name string) (*entity.ConfectionaryType, error)
	List(limit, offset int) ([]*entity.ConfectionaryType, error)
	Update(city *entity.ConfectionaryType) error
	Delete(id int) error
}

func (s *server) CreateConfectionaryType(ctx context.Context, request *api.CreateConfectionaryTypeRequest) (*api.ConfectionaryType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByIdConfectionaryType(ctx context.Context, request *api.FindByIdRequest) (*api.ConfectionaryType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByNameConfectionaryType(ctx context.Context, request *api.FindByNameRequest) (*api.ConfectionaryType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListConfectionaryType(ctx context.Context, request *api.ListRequest) (*api.ListConfectionaryTypeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateConfectionaryType(ctx context.Context, city *api.ConfectionaryType) (*api.ConfectionaryType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeleteConfectionaryType(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}
