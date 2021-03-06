package kratos

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
)

// Option is an application option.
type Option func(o *options)

// options is an application options.
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []string

	ctx  context.Context
	sigs []os.Signal

	logger   log.Logger
	registry registry.Registry
	servers  []transport.Server
}

func (o options) Service() *registry.Service {
	return &registry.Service{
		ID:        o.id,
		Name:      o.name,
		Version:   o.version,
		Metadata:  o.metadata,
		Endpoints: o.endpoints,
	}
}

// ID with service id.
func ID(id string) Option {
	return func(o *options) { o.id = id }
}

// Name with service name.
func Name(name string) Option {
	return func(o *options) { o.name = name }
}

// Version with service version.
func Version(version string) Option {
	return func(o *options) { o.version = version }
}

// Metadata with service metadata.
func Metadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// Endpoint with service endpoint.
func Endpoint(endpoints ...string) Option {
	return func(o *options) { o.endpoints = endpoints }
}

// Context with service context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Signal with exit signals.
func Signal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// Logger with service logger.
func Logger(logger log.Logger) Option {
	return func(o *options) { o.logger = logger }
}

// Registry with service registry.
func Registry(r registry.Registry) Option {
	return func(o *options) { o.registry = r }
}

// Server with transport servers.
func Server(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}
