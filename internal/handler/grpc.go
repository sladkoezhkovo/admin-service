package handler

import "github.com/sladkoezhkovo/admin-service/api"

// var _ (api.AdminServiceServer) = (*server)(nil)

var (
	EMPTY_RESPONSE = &api.Empty{}
)

type server struct {
	api.UnimplementedAdminServiceServer

	city              CityService
	district          DistrictService
	packaging         PackagingService
	unit              UnitService
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
		city:              cityService,
		district:          districtService,
		packaging:         packagingService,
		unit:              unitService,
		propertyType:      propertyTypeService,
		confectionaryType: confectionaryTypeService,
	}
}
