package handler

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type DistrictService interface {
	Create(city *entity.District) error
	FindById(id int) (*entity.District, error)
	FindByName(name string) (*entity.District, error)
	List(limit, offset int) ([]*entity.District, error)
	Update(city *entity.District) error
	Delete(id int) error
}

func (s *server) CreateDistrict(ctx context.Context, request *api.CreateDistrictRequest) (*api.District, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByIdDistrict(ctx context.Context, request *api.FindByIdRequest) (*api.District, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByNameDistrict(ctx context.Context, request *api.FindByNameRequest) (*api.District, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListDistrict(ctx context.Context, request *api.ListRequest) (*api.ListDistrictResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateDistrict(ctx context.Context, city *api.District) (*api.District, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeleteDistrict(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}
