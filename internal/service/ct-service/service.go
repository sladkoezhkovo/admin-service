package confectionaryservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/service"
)

type ConfectionaryTypeRepository interface {
	Create(confectionary *entity.ConfectionaryType) error
	FindById(id int64) (*entity.ConfectionaryType, error)
	List(limit, offset int) ([]*entity.ConfectionaryType, int64, error)
	Update(confectionary *entity.ConfectionaryType) error
	Delete(id int64) error
}

var _ handler.ConfectionaryTypeService = (*confectionaryService)(nil)

type confectionaryService struct {
	repository ConfectionaryTypeRepository
}

func New(repository ConfectionaryTypeRepository) *confectionaryService {
	return &confectionaryService{
		repository: repository,
	}
}

func (c *confectionaryService) Create(confectionary *entity.ConfectionaryType) error {
	if err := c.repository.Create(confectionary); err != nil {
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

func (c *confectionaryService) FindById(id int64) (*entity.ConfectionaryType, error) {
	confectionary, err := c.repository.FindById(id)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return confectionary, nil
}

func (c *confectionaryService) List(limit, offset int) ([]*entity.ConfectionaryType, int64, error) {
	if !(limit > 0) {
		return nil, 0, service.ErrInvalidLimit
	} else if offset < 0 {
		return nil, 0, service.ErrInvalidOffset
	}

	ctct, count, err := c.repository.List(limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, 0, pgerr
		}
		return nil, 0, err
	}
	return ctct, count, nil
}

func (c *confectionaryService) Update(confectionary *entity.ConfectionaryType) error {
	if err := c.repository.Update(confectionary); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}

func (c *confectionaryService) Delete(id int64) error {
	if err := c.repository.Delete(id); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}
