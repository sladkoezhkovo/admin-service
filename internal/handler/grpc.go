package handler

import "github.com/sladkoezhkovo/admin-service/api"

// var _ (api.AdminServiceServer) = (*server)(nil)

type server struct {
	api.UnimplementedAdminServiceServer

	cities            CityService
	districts         DistrictService
	packaging         PackagingService
	units             UnitService
	propertyType      PropertyTypeService
	confectionaryType ConfectionaryTypeService
}
