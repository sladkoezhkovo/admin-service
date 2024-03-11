package handler

import (
	"context"
	"fmt"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/converter"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type DistrictService interface {
	Create(district *entity.District) error
	FindById(id int64) (*entity.District, error)
	FindByName(name string) (*entity.District, error)
	List(limit, offset int) ([]*entity.District, error)
	Update(district *entity.District) error
	Delete(id int64) error
}

func (s *server) CreateDistrict(ctx context.Context, request *api.CreateDistrictRequest) (*api.District, error) {

	district := converter.DistrictCreateDtoToEntity(request)

	city, err := s.city.FindById(request.CityId)
	if err != nil {
		return nil, err
	}

	if err := s.district.Create(district); err != nil {
		return nil, err
	}

	fmt.Printf("created district: %v\n", district)

	response := &api.District{
		Id:   district.Id,
		Name: district.Name,
		City: city.Name,
	}

	return response, nil
}

func (s *server) FindByIdDistrict(ctx context.Context, request *api.FindByIdRequest) (*api.District, error) {
	district, err := s.district.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	response := converter.DistrictEntityToDto(district)

	return response, nil
}

func (s *server) FindByNameDistrict(ctx context.Context, request *api.FindByNameRequest) (*api.District, error) {
	district, err := s.district.FindByName(request.Name)
	if err != nil {
		return nil, err
	}

	return &api.District{
		Id:   district.Id,
		Name: district.Name,
	}, nil
}

func (s *server) ListDistrict(ctx context.Context, request *api.ListRequest) (*api.ListDistrictResponse, error) {
	districts, err := s.district.List(int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	response := &api.ListDistrictResponse{
		Entries: make([]*api.District, 0, len(districts)),
	}

	for _, district := range districts {
		response.Entries = append(response.Entries, &api.District{
			Id:   district.Id,
			Name: district.Name,
			City: district.City.Name,
		})
	}

	return response, nil
}

func (s *server) UpdateDistrict(ctx context.Context, request *api.UpdateDistrictRequest) (*api.District, error) {
	district := &entity.District{
		Id:   request.Id,
		Name: request.Name,
		City: entity.City{
			Id: request.CityId,
		},
	}

	if err := s.district.Update(district); err != nil {
		return nil, err
	}

	city, err := s.city.FindById(request.CityId)
	if err != nil {
		return nil, err
	}

	response := &api.District{
		Id:   district.Id,
		Name: district.Name,
		City: city.Name,
	}

	return response, nil
}

func (s *server) DeleteDistrict(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	if err := s.district.Delete(request.Id); err != nil {
		return nil, err
	}
	return EMPTY_RESPONSE, nil
}
