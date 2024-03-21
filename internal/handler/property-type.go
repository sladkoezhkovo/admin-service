package handler

import (
	"context"
	api "github.com/sladkoezhkovo/admin-service/api/admin"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type PropertyTypeService interface {
	Create(propertyType *entity.PropertyType) error
	FindById(id int64) (*entity.PropertyType, error)
	List(limit, offset int) ([]*entity.PropertyType, int64, error)
	Update(propertyType *entity.PropertyType) error
	Delete(id int64) error
}

func (s *server) CreatePropertyType(ctx context.Context, request *api.CreatePropertyTypeRequest) (*api.PropertyType, error) {
	propertyType := &entity.PropertyType{
		Name: request.Name,
	}

	if err := s.propertyType.Create(propertyType); err != nil {
		return nil, err
	}

	return &api.PropertyType{
		Id:   propertyType.Id,
		Name: propertyType.Name,
	}, nil
}

func (s *server) FindByIdPropertyType(ctx context.Context, request *api.FindByIdRequest) (*api.PropertyType, error) {
	propertyType, err := s.propertyType.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &api.PropertyType{
		Id:   propertyType.Id,
		Name: propertyType.Name,
	}, nil
}

func (s *server) ListPropertyType(ctx context.Context, request *api.Bounds) (*api.ListPropertyTypeResponse, error) {
	ptpt, count, err := s.propertyType.List(int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	response := &api.ListPropertyTypeResponse{
		Entries: make([]*api.PropertyType, 0, len(ptpt)),
		Count:   count,
	}

	for _, propertyType := range ptpt {
		response.Entries = append(response.Entries, &api.PropertyType{
			Id:   propertyType.Id,
			Name: propertyType.Name,
		})
	}

	return response, nil
}

func (s *server) UpdatePropertyType(ctx context.Context, request *api.PropertyType) (*api.PropertyType, error) {
	propertyType := &entity.PropertyType{
		Id:   request.Id,
		Name: request.Name,
	}

	if err := s.propertyType.Update(propertyType); err != nil {
		return nil, err
	}

	return request, nil
}

func (s *server) DeletePropertyType(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	if err := s.propertyType.Delete(request.Id); err != nil {
		return nil, err
	}
	return EMPTY_RESPONSE, nil
}
