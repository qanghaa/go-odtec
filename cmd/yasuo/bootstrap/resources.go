package bootstrap

import (
	"go-odtec/utils/database"

	"go.uber.org/zap"
)

// Resources contains the platform service clients which can be
// used by the server.
type Resources struct {
	logger *zap.Logger
	db     *database.DBTrace
}

type ResourcesOption func(resources *Resources)

// NewResources creates a new Resources.
func NewResources(opts ...ResourcesOption) *Resources {
	r := &Resources{}
	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithLogger(logger *zap.Logger) ResourcesOption {
	return ResourcesOption(func(r *Resources) {
		r.logger = logger
	})
}

func WithDB(db *database.DBTrace) ResourcesOption {
	return ResourcesOption(func(r *Resources) {
		r.db = db
	})
}

func (r *Resources) Logger() *zap.Logger {
	if r.logger == nil {
		panic("logger not initialized")
	}
	return r.logger
}

func (r *Resources) DB() *database.DBTrace {
	if r.db == nil {
		panic("database not initialized")
	}
	return r.db
}
