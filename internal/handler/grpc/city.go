package grpc

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type CityService interface {
	Create(city *entity.City) error
	FindById(id int) (*entity.City, error)
	FindByName(name string) (*entity.City, error)
	List(limit, offset int) ([]*entity.City, error)
	Update(city *entity.City) error
	Delete(id int) error
}

func (s *server) CreateCity(ctx context.Context, request *api.CreateCityRequest) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByIdCity(ctx context.Context, request *api.FindByIdRequest) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByNameCity(ctx context.Context, request *api.FindByNameRequest) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListCity(ctx context.Context, request *api.ListRequest) (*api.ListCityResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateCity(ctx context.Context, city *api.City) (*api.City, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeleteCity(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}
