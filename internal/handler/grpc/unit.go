package grpc

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type UnitService interface {
	Create(city *entity.Unit) error
	FindById(id int) (*entity.Unit, error)
	FindByName(name string) (*entity.Unit, error)
	List(limit, offset int) ([]*entity.Unit, error)
	Update(city *entity.Unit) error
	Delete(id int) error
}

func (s *server) CreateUnit(ctx context.Context, request *api.CreateCityRequest) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByIdUnit(ctx context.Context, request *api.FindByIdRequest) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByNameUnit(ctx context.Context, request *api.FindByNameRequest) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListUnit(ctx context.Context, request *api.ListRequest) (*api.ListCityResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateUnit(ctx context.Context, city *api.City) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeleteUnit(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}
