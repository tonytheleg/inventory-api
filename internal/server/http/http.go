package http

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// New create a new http server.
func New(c CompletedConfig, authn middleware.Middleware) *http.Server {
	// TODO: pass in health, authn middleware
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			authn,
		),
	}
	opts = append(opts, c.ServerOptions...)
	srv := http.NewServer(opts...)
	return srv
}