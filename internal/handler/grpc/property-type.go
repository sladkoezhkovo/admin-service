package grpc

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type PropertyTypeService interface {
	Create(city *entity.PropertyType) error
	FindById(id int) (*entity.PropertyType, error)
	FindByName(name string) (*entity.PropertyType, error)
	List(limit, offset int) ([]*entity.PropertyType, error)
	Update(city *entity.PropertyType) error
	Delete(id int) error
}

func (s *server) CreatePropertyType(ctx context.Context, request *api.CreatePropertyTypeRequest) (*api.PropertyType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByIdPropertyType(ctx context.Context, request *api.FindByIdRequest) (*api.PropertyType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByNamePropertyType(ctx context.Context, request *api.FindByNameRequest) (*api.PropertyType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListPropertyType(ctx context.Context, request *api.ListRequest) (*api.ListPropertyTypeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdatePropertyType(ctx context.Context, city *api.PropertyType) (*api.PropertyType, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeletePropertyType(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}
