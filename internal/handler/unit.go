package handler

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type UnitService interface {
	Create(unit *entity.Unit) error
	FindById(id int64) (*entity.Unit, error)
	ListByName(name string, limit, offset int32) ([]*entity.Unit, error)
	List(limit, offset int) ([]*entity.Unit, error)
	Update(unit *entity.Unit) error
	Delete(id int64) error
}

func (s *server) CreateUnit(ctx context.Context, request *api.CreateUnitRequest) (*api.Unit, error) {
	unit := &entity.Unit{
		Name: request.Name,
	}

	if err := s.unit.Create(unit); err != nil {
		return nil, err
	}

	return &api.Unit{
		Id:   unit.Id,
		Name: unit.Name,
	}, nil
}

func (s *server) FindByIdUnit(ctx context.Context, request *api.FindByIdRequest) (*api.Unit, error) {
	unit, err := s.unit.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &api.Unit{
		Id:   unit.Id,
		Name: unit.Name,
	}, nil
}

func (s *server) ListByNameUnit(ctx context.Context, request *api.ListByNameRequest) (*api.ListUnitResponse, error) {
	entries, err := s.unit.ListByName(request.Name, request.Bounds.Limit, request.Bounds.Offset)
	if err != nil {
		return nil, err
	}

	response := &api.ListUnitResponse{
		Entries: make([]*api.Unit, 0, len(entries)),
	}

	for _, e := range entries {
		response.Entries = append(response.Entries, &api.Unit{
			Id:   e.Id,
			Name: e.Name,
		})

	}
	return response, nil
}

func (s *server) ListUnit(ctx context.Context, request *api.ListRequest) (*api.ListUnitResponse, error) {
	units, err := s.unit.List(int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	response := &api.ListUnitResponse{
		Entries: make([]*api.Unit, 0, len(units)),
	}

	for _, unit := range units {
		response.Entries = append(response.Entries, &api.Unit{
			Id:   unit.Id,
			Name: unit.Name,
		})
	}

	return response, nil
}

func (s *server) UpdateUnit(ctx context.Context, request *api.Unit) (*api.Unit, error) {
	unit := &entity.Unit{
		Id:   request.Id,
		Name: request.Name,
	}

	if err := s.unit.Update(unit); err != nil {
		return nil, err
	}

	return request, nil
}

func (s *server) DeleteUnit(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	if err := s.unit.Delete(request.Id); err != nil {
		return nil, err
	}
	return EMPTY_RESPONSE, nil
}
