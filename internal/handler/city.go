package handler

import (
	"context"
	"fmt"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type CityService interface {
	Create(city *entity.City) error
	FindById(id int64) (*entity.City, error)
	FindByName(name string) (*entity.City, error)
	List(limit, offset int) ([]*entity.City, error)
	Update(city *entity.City) error
	Delete(id int64) error
}

func (s *server) CreateCity(ctx context.Context, request *api.CreateCityRequest) (*api.City, error) {
	city := &entity.City{
		Name: request.Name,
	}

	if err := s.city.Create(city); err != nil {
		return nil, err
	}

	fmt.Printf("created city: %v\n", city)

	return &api.City{
		Id:   city.Id,
		Name: city.Name,
	}, nil
}

func (s *server) FindByIdCity(ctx context.Context, request *api.FindByIdRequest) (*api.City, error) {
	city, err := s.city.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &api.City{
		Id:   city.Id,
		Name: city.Name,
	}, nil
}

func (s *server) FindByNameCity(ctx context.Context, request *api.FindByNameRequest) (*api.City, error) {
	city, err := s.city.FindByName(request.Name)
	if err != nil {
		return nil, err
	}

	return &api.City{
		Id:   city.Id,
		Name: city.Name,
	}, nil
}

func (s *server) ListCity(ctx context.Context, request *api.ListRequest) (*api.ListCityResponse, error) {
	cities, err := s.city.List(int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	response := &api.ListCityResponse{
		Entries: make([]*api.City, 0, len(cities)),
	}

	for _, city := range cities {
		response.Entries = append(response.Entries, &api.City{
			Id:   city.Id,
			Name: city.Name,
		})
	}

	return response, nil
}

func (s *server) UpdateCity(ctx context.Context, request *api.City) (*api.City, error) {
	city := &entity.City{
		Id:   request.Id,
		Name: request.Name,
	}

	if err := s.city.Update(city); err != nil {
		return nil, err
	}

	return request, nil
}

func (s *server) DeleteCity(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	if err := s.city.Delete(request.Id); err != nil {
		return nil, err
	}
	return EMPTY_RESPONSE, nil
}
