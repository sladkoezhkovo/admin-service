package handler

import "github.com/sladkoezhkovo/admin-service/api"

// var _ (api.AdminServiceServer) = (*server)(nil)

var (
	EMPTY_RESPONSE = &api.Empty{}
)

type server struct {
	api.UnimplementedAdminServiceServer

	cities            CityService
	districts         DistrictService
	packaging         PackagingService
	units             UnitService
	propertyType      PropertyTypeService
	confectionaryType ConfectionaryTypeService
}

func New(
	cityService CityService,
	districtService DistrictService,
	packagingService PackagingService,
	unitService UnitService,
	propertyTypeService PropertyTypeService,
	confectionaryTypeService ConfectionaryTypeService,
) *server {
	return &server{
		cities:            cityService,
		districts:         districtService,
		packaging:         packagingService,
		units:             unitService,
		propertyType:      propertyTypeService,
		confectionaryType: confectionaryTypeService,
	}
}
