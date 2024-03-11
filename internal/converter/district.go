package converter

import (
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
)

func DistrictCreateDtoToEntity(dto *api.CreateDistrictRequest) *entity.District {
	return &entity.District{
		Name: dto.Name,
		City: entity.City{
			Id: dto.CityId,
		},
	}
}

func DistrictEntityToDto(e *entity.District) *api.District {
	return &api.District{
		Id:   e.Id,
		Name: e.Name,
		City: e.City.Name,
	}
}
