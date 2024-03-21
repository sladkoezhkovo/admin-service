package handler

import (
	"context"
	api "github.com/sladkoezhkovo/admin-service/api/admin"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type PackagingService interface {
	Create(packaging *entity.Packaging) error
	FindById(id int64) (*entity.Packaging, error)
	List(limit, offset int) ([]*entity.Packaging, int64, error)
	Update(packaging *entity.Packaging) error
	Delete(id int64) error
}

func (s *server) CreatePackaging(ctx context.Context, request *api.CreatePackagingRequest) (*api.Packaging, error) {
	packaging := &entity.Packaging{
		Name: request.Name,
	}

	if err := s.packaging.Create(packaging); err != nil {
		return nil, err
	}

	return &api.Packaging{
		Id:   packaging.Id,
		Name: packaging.Name,
	}, nil
}

func (s *server) FindByIdPackaging(ctx context.Context, request *api.FindByIdRequest) (*api.Packaging, error) {
	packaging, err := s.packaging.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	return &api.Packaging{
		Id:   packaging.Id,
		Name: packaging.Name,
	}, nil
}

func (s *server) ListPackaging(ctx context.Context, request *api.Bounds) (*api.ListPackagingResponse, error) {
	pp, count, err := s.packaging.List(int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	response := &api.ListPackagingResponse{
		Entries: make([]*api.Packaging, 0, len(pp)),
		Count:   count,
	}

	for _, p := range pp {
		response.Entries = append(response.Entries, &api.Packaging{
			Id:   p.Id,
			Name: p.Name,
		})
	}

	return response, nil
}

func (s *server) UpdatePackaging(ctx context.Context, request *api.Packaging) (*api.Packaging, error) {
	packaging := &entity.Packaging{
		Id:   request.Id,
		Name: request.Name,
	}

	if err := s.packaging.Update(packaging); err != nil {
		return nil, err
	}

	return request, nil
}

func (s *server) DeletePackaging(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	if err := s.packaging.Delete(request.Id); err != nil {
		return nil, err
	}
	return EMPTY_RESPONSE, nil
}
