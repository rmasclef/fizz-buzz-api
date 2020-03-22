package di

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/gol4ng/httpware/v2"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger-http/middleware"
	"github.com/gorilla/mux"

	"github.com/rmasclef/fizz_buzz_api/internal/http/handler"
	internal_middleware "github.com/rmasclef/fizz_buzz_api/internal/http/middleware"
	"github.com/rmasclef/fizz_buzz_api/internal/http/server"
	"github.com/rmasclef/fizz_buzz_api/pkg/fizz-buzz"
)

func (container *Container) GetHTTPServer(ctx context.Context) *server.Server {
	if container.httpServer == nil {
		stack := container.getHTTPMiddlewares()

		srv := &http.Server{
			Addr:              container.Cfg.HTTPAddr,
			Handler:           stack.DecorateHandler(container.getHTTPHandler()),
			ReadHeaderTimeout: container.Cfg.ReadHeaderTimeout,
			WriteTimeout:      container.Cfg.WriteTimeout,
			IdleTimeout:       container.Cfg.IdleTimeout,
			MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
			BaseContext: func(_ net.Listener) context.Context {
				return logger.InjectInContext(ctx, container.GetLogger())
			},
		}
		container.httpServer = &server.Server{Srv: srv, Log: container.GetLogger()}
	}

	return container.httpServer
}

func (container *Container) getHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/fizz-buzz", container.getFizzbuzzHandler()).Methods(http.MethodPost)

	// enable profiling with pprof if DEBUG is asked
	if container.Cfg.Debug {
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/allocs", pprof.Index)
		router.HandleFunc("/debug/pprof/block", pprof.Index)
		router.HandleFunc("/debug/pprof/heap", pprof.Index)
		router.HandleFunc("/debug/pprof/goroutine", pprof.Index)
		router.HandleFunc("/debug/pprof/mutex", pprof.Index)
		router.HandleFunc("/debug/pprof/threadcreate", pprof.Index)

		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	return router
}

func (container *Container) getHTTPMiddlewares() httpware.Middlewares {
	return httpware.MiddlewareStack(
		middleware.CorrelationId(),
		middleware.Logger(container.GetLogger()),
		internal_middleware.ContentTypeFilterer,
	)
}

func (container *Container) getFizzbuzzHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		/*
			/!\ At this time we only support JSON requests and responses so we use the JSONTransformer ...
			... and we always set the content-type to application/json

			if we want to accept other format of request bodies
			we should create a middleware that defines the right transformer
			using the request content-type and accept header
		*/

		handler.FizzBuzz(
			&fizzbuzz.JSONTransformer{},
			fizzbuzz.RequestValidator,
			fizzbuzz.FizzBuzzController,
		).ServeHTTP(writer, request)

		writer.Header().Set("Content-Type", "application/json")
	}
}
