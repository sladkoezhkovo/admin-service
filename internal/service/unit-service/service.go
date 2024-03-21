package unitservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/service"
)

type UnitRepository interface {
	Create(unit *entity.Unit) error
	FindById(id int64) (*entity.Unit, error)
	List(limit, offset int) ([]*entity.Unit, int64, error)
	Update(unit *entity.Unit) error
	Delete(id int64) error
}

var _ handler.UnitService = (*unitService)(nil)

type unitService struct {
	repository UnitRepository
}

func New(repository UnitRepository) *unitService {
	return &unitService{
		repository: repository,
	}
}

func (c *unitService) Create(unit *entity.Unit) error {
	if err := c.repository.Create(unit); err != nil {
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

func (c *unitService) FindById(id int64) (*entity.Unit, error) {
	unit, err := c.repository.FindById(id)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return unit, nil
}

func (c *unitService) List(limit, offset int) ([]*entity.Unit, int64, error) {
	if !(limit > 0) {
		return nil, 0, service.ErrInvalidLimit
	} else if offset < 0 {
		return nil, 0, service.ErrInvalidOffset
	}

	uu, count, err := c.repository.List(limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, 0, pgerr
		}
		return nil, 0, err
	}
	return uu, count, nil
}

func (c *unitService) Update(unit *entity.Unit) error {
	if err := c.repository.Update(unit); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}

func (c *unitService) Delete(id int64) error {
	if err := c.repository.Delete(id); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}
