package abstraction

import "context"

type Server interface {
	ServerName() string
	InitDependencies()
	GratefulShutdown(ctx context.Context) error
}
