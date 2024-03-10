package grpc

import (
	"context"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

type PackagingService interface {
	Create(city *entity.Packaging) error
	FindById(id int) (*entity.Packaging, error)
	FindByName(name string) (*entity.Packaging, error)
	List(limit, offset int) ([]*entity.Packaging, error)
	Update(city *entity.Packaging) error
	Delete(id int) error
}

func (s *server) CreatePackaging(ctx context.Context, request *api.CreatePackagingRequest) (*api.Packaging, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByIdPackaging(ctx context.Context, request *api.FindByIdRequest) (*api.Packaging, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) FindByNamePackaging(ctx context.Context, request *api.FindByNameRequest) (*api.Packaging, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListPackaging(ctx context.Context, request *api.ListRequest) (*api.ListPackagingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdatePackaging(ctx context.Context, city *api.Packaging) (*api.Packaging, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeletePackaging(ctx context.Context, request *api.FindByIdRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}
