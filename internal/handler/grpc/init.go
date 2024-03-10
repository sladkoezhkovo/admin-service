package grpc

import "github.com/sladkoezhkovo/admin-service/api"

// var _ (api.AdminServiceServer) = (*server)(nil)

type server struct {
	api.UnimplementedAdminServiceServer
}
