package abstraction

import "google.golang.org/grpc"

type GRPCService interface {
	Register(*grpc.Server)
}
