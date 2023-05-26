package yasuo

import (
	"go-odtec/cmd/yasuo/bootstrap"
	"go-odtec/cmd/yasuo/config"
	repo "go-odtec/domain/yasuo/repository"
	"go-odtec/domain/yasuo/server"
	srv "go-odtec/domain/yasuo/service"
	gtrans "go-odtec/domain/yasuo/transport/grpc"
	"go-odtec/utils/database"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
)

func InitServices(db database.QueryExecer) (srvs []server.GRPCService) {
	srvs = append(srvs, gtrans.NewUserGRPC(srv.NewUserService(&repo.UserRepo{}, db)))
	// register others here
	return
}

func unaryServerInterceptors(rsc *bootstrap.Resources) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		grpc_zap.UnaryServerInterceptor(rsc.Logger(), grpc_zap.WithLevels(grpc_zap.DefaultCodeToLevel)),
		// add another interceptors here
	)
}

func StartGRPCServer(rsc *bootstrap.Resources, cfg *config.Config) *server.Server {
	opts := unaryServerInterceptors(rsc)
	server := server.New(opts)
	srvs := InitServices(rsc.DB())
	server.WithRegisteredServices(srvs...)
	return server
}
