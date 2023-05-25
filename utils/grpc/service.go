package grpc

import "google.golang.org/grpc"

type GRPCService interface {
	Register(*grpc.Server)
}
