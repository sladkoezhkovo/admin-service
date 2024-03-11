package districtservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/service"
	"strings"
)

type DistrictRepository interface {
	Create(district *entity.District) error
	FindById(id int64) (*entity.District, error)
	ListByCityId(cityId int64) ([]*entity.District, error)
	ListByName(name string, limit, offset int32) ([]*entity.District, error)
	List(limit, offset int) ([]*entity.District, error)
	Update(district *entity.District) error
	Delete(id int64) error
}

var _ handler.DistrictService = (*districtService)(nil)

type districtService struct {
	repository DistrictRepository
}

func New(repository DistrictRepository) *districtService {
	return &districtService{
		repository: repository,
	}
}

func (c *districtService) Create(district *entity.District) error {

	if dd, err := c.repository.ListByCityId(district.City.Id); err == nil {
		for _, d := range dd {
			if strings.Compare(strings.ToLower(d.Name), strings.ToLower(district.Name)) == 0 {
				return service.ErrUniqueViolation
			}
		}
	}

	if err := c.repository.Create(district); err != nil {
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

func (c *districtService) FindById(id int64) (*entity.District, error) {
	district, err := c.repository.FindById(id)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return district, nil
}

func (c *districtService) ListByName(name string, limit, offset int32) ([]*entity.District, error) {
	dd, err := c.repository.ListByName(name, limit, offset)
	if err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return nil, pgerr
		}
		return nil, err
	}
	return dd, nil
}

func (c *districtService) List(limit, offset int) ([]*entity.District, error) {
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

func (c *districtService) Update(district *entity.District) error {
	if err := c.repository.Update(district); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}

func (c *districtService) Delete(id int64) error {
	if err := c.repository.Delete(id); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			return pgerr
		}
		return err
	}
	return nil
}
