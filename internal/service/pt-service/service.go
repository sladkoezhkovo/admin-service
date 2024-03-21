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
	List(limit, offset int) ([]*entity.PropertyType, int64, error)
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

func (c *ptService) List(limit, offset int) ([]*entity.PropertyType, int64, error) {
	if !(limit > 0) {
		return nil, 0, service.ErrInvalidLimit
	} else if offset < 0 {
		return nil, 0, service.ErrInvalidOffset
	}

	ptpt, count, err := c.repository.List(limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, 0, pgerr
		}
		return nil, 0, err
	}
	return ptpt, count, nil
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
