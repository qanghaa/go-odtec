package server

import (
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	services   []GRPCService
}

type GRPCService interface {
	Register(*grpc.Server)
}

func (s *Server) ServerName() string {
	return "yasuo"
}

func (s *Server) WithRegisteredServices(services ...GRPCService) {
	if s.services == nil {
		s.services = []GRPCService{}
	}
	s.services = append(s.services, services...)
	for _, srv := range services {
		srv.Register(s.grpcServer)
	}
}

func New(opts ...grpc.ServerOption) *Server {
	var (
		grpcServer = grpc.NewServer(opts...)
		s          = &Server{
			grpcServer: grpcServer,
		}
	)
	return s
}

func (s *Server) InitDependencies() {

}

func (s *Server) GratefulShutdown() {

}
