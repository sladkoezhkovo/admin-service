package ptservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/service"
)

type PropertyTypeRepository interface {
	Create(pt *entity.PropertyType) error
	FindById(id int64) (*entity.PropertyType, error)
	ListByName(name string, limit, offset int32) ([]*entity.PropertyType, error)
	List(limit, offset int) ([]*entity.PropertyType, error)
	Update(pt *entity.PropertyType) error
	Delete(id int64) error
}

var _ handler.PropertyTypeService = (*ptService)(nil)

type ptService struct {
	repository PropertyTypeRepository
}

func New(repository PropertyTypeRepository) *ptService {
	return &ptService{
		repository: repository,
	}
}

func (c *ptService) Create(pt *entity.PropertyType) error {
	if err := c.repository.Create(pt); err != nil {
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

func (c *ptService) FindById(id int64) (*entity.PropertyType, error) {
	pt, err := c.repository.FindById(id)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return pt, nil
}

func (c *ptService) ListByName(name string, limit, offset int32) ([]*entity.PropertyType, error) {
	pt, err := c.repository.ListByName(name, limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return pt, nil
}

func (c *ptService) List(limit, offset int) ([]*entity.PropertyType, error) {
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

func (c *ptService) Update(pt *entity.PropertyType) error {
	if err := c.repository.Update(pt); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}

func (c *ptService) Delete(id int64) error {
	if err := c.repository.Delete(id); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}
