package server

import (
	repo "go-odtec/domain/yasuo/repository"
	srv "go-odtec/domain/yasuo/service"
	gtrans "go-odtec/domain/yasuo/transport/grpc"
	"go-odtec/utils/database"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	services   []gRPCService
	db         database.QueryExecer
}

type gRPCService interface {
	Register(*grpc.Server)
}

func (s *Server) ServerName() string {
	return "yasuo"
}

func (s *Server) WithRegisteredServices(services ...gRPCService) {
	if s.services == nil {
		s.services = []gRPCService{}
	}
	s.services = append(s.services, services...)
	for _, srv := range services {
		srv.Register(s.grpcServer)
	}
}

func New(db database.QueryExecer, opts ...grpc.ServerOption) *Server {
	var (
		grpcServer = grpc.NewServer(opts...)
		s          = &Server{
			grpcServer: grpcServer,
			db:         db,
		}
	)
	return s
}

func InitServices(db database.QueryExecer) (srvs []gRPCService) {
	srvs = append(srvs, gtrans.NewUserGRPC(srv.NewUserService(&repo.UserRepo{}, db)))
	return
}