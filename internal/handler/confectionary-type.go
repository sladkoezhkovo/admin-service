package handler

import (
	"context"
	api "github.com/sladkoezhkovo/admin-service/api/admin"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type ConfectionaryTypeService interface {
	Create(confectionaryType *entity.ConfectionaryType) error
	FindById(id int64) (*entity.ConfectionaryType, error)
	List(limit, offset int) ([]*entity.ConfectionaryType, int64, error)
	Update(confectionaryType *entity.ConfectionaryType) error
	Delete(id int64) error
}

func (s *server) CreateConfectionaryType(ctx context.Context, request *api.CreateConfectionaryTypeRequest) (*api.ConfectionaryType, error) {
	confectionaryType := &entity.ConfectionaryType{
		Name: request.Name,
	}

	if err := s.confectionaryType.Create(confectionaryType); err != nil {
		return nil, err
	}

	return &api.ConfectionaryType{
		Id:   confectionaryType.Id,
		Name: confectionaryType.Name,
	}, nil
}

func (s *server) FindByIdConfectionaryType(ctx context.Context, request *api.FindByIdRequest) (*api.ConfectionaryType, error) {
	confectionaryType, err := s.confectionaryType.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &api.ConfectionaryType{
		Id:   confectionaryType.Id,
		Name: confectionaryType.Name,
	}, nil
}

func (s *server) ListConfectionaryType(ctx context.Context, request *api.Bounds) (*api.ListConfectionaryTypeResponse, error) {
	cts, count, err := s.confectionaryType.List(int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	response := &api.ListConfectionaryTypeResponse{
		Entries: make([]*api.ConfectionaryType, 0, len(cts)),
		Count:   count,
	}

	for _, confectionaryType := range cts {
		response.Entries = append(response.Entries, &api.ConfectionaryType{
			Id:   confectionaryType.Id,
			Name: confectionaryType.Name,
		})
	}

	return response, nil
}

func (s *server) UpdateConfectionaryType(ctx context.Context, request *api.ConfectionaryType) (*api.ConfectionaryType, error) {
	confectionaryType := &entity.ConfectionaryType{
		Id:   request.Id,
		Name: request.Name,
	}

	if err := s.confectionaryType.Update(confectionaryType); err != nil {
		return nil, err
	}

	return request, nil
}

func (s *server) DeleteConfectionaryType(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	if err := s.confectionaryType.Delete(request.Id); err != nil {
		return nil, err
	}
	return EMPTY_RESPONSE, nil
}
