package cityservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/service"
)

type CityRepository interface {
	Create(city *entity.City) error
	FindById(id int64) (*entity.City, error)
	ListByName(name string, limit, offset int32) ([]*entity.City, error)
	List(limit, offset int) ([]*entity.City, error)
	Update(city *entity.City) error
	Delete(id int64) error
}

var _ handler.CityService = (*cityService)(nil)

type cityService struct {
	repository CityRepository
}

func New(repository CityRepository) *cityService {
	return &cityService{
		repository: repository,
	}
}

func (c *cityService) Create(city *entity.City) error {
	if err := c.repository.Create(city); err != nil {
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

func (c *cityService) FindById(id int64) (*entity.City, error) {
	city, err := c.repository.FindById(id)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return city, nil
}

func (c *cityService) ListByName(name string, limit, offset int32) ([]*entity.City, error) {
	city, err := c.repository.ListByName(name, limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return city, nil
}

func (c *cityService) List(limit, offset int) ([]*entity.City, error) {
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

func (c *cityService) Update(city *entity.City) error {
	if err := c.repository.Update(city); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}

func (c *cityService) Delete(id int64) error {
	if err := c.repository.Delete(id); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}
