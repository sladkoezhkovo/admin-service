package packagingservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/service"
)

type PackagingRepository interface {
	Create(packaging *entity.Packaging) error
	FindById(id int64) (*entity.Packaging, error)
	ListByName(name string, limit, offset int32) ([]*entity.Packaging, error)
	List(limit, offset int) ([]*entity.Packaging, error)
	Update(packaging *entity.Packaging) error
	Delete(id int64) error
}

var _ handler.PackagingService = (*packagingService)(nil)

type packagingService struct {
	repository PackagingRepository
}

func New(repository PackagingRepository) *packagingService {
	return &packagingService{
		repository: repository,
	}
}

func (c *packagingService) Create(packaging *entity.Packaging) error {
	if err := c.repository.Create(packaging); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			switch pgerr.Code {
			case "23505":
				return service.ErrUniqueViolation
			}
			return pgerr
		}
		return err
	}
	return nil
}

func (c *packagingService) FindById(id int64) (*entity.Packaging, error) {
	packaging, err := c.repository.FindById(id)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return packaging, nil
}

func (c *packagingService) ListByName(name string, limit, offset int32) ([]*entity.Packaging, error) {
	packaging, err := c.repository.ListByName(name, limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return packaging, nil
}

func (c *packagingService) List(limit, offset int) ([]*entity.Packaging, error) {
	if !(limit > 0) {
		return nil, service.ErrInvalidLimit
	} else if offset < 0 {
		return nil, service.ErrInvalidOffset
	}

	cities, err := c.repository.List(limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return cities, nil
}

func (c *packagingService) Update(packaging *entity.Packaging) error {
	if err := c.repository.Update(packaging); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}

func (c *packagingService) Delete(id int64) error {
	if err := c.repository.Delete(id); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}
