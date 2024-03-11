package converter

import (
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

func CityDtoToEntity(dto *api.City) *entity.City {
	return &entity.City{
		Id:   dto.Id,
		Name: dto.Name,
	}
}

func CityEntityToDto(e *entity.City) *api.City {
	return &api.City{
		Id:   e.Id,
		Name: e.Name,
	}
}
