package handler

import (
	api "github.com/sladkoezhkovo/admin-service/api/admin"
	"log/slog"
	"os"
)

// var _ (api.AdminServiceServer) = (*server)(nil)

var (
	EMPTY_RESPONSE = &api.Empty{}
)

type server struct {
	api.UnimplementedAdminServiceServer

	logger *slog.Logger

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
		logger:            slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		city:              cityService,
		district:          districtService,
		packaging:         packagingService,
		unit:              unitService,
		propertyType:      propertyTypeService,
		confectionaryType: confectionaryTypeService,
	}
}
