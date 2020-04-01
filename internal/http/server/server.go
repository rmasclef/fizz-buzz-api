package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gol4ng/logger"
)

type Server struct {
	Srv *http.Server
	Log logger.LoggerInterface
}

func (s *Server) Start(stopChan chan<- interface{}) {
	go func() {
		s.Log.Info("starting HTTP server...", logger.String("addr", s.Srv.Addr))
		if err := s.Srv.ListenAndServe(); err != nil {
			stopChan <- fmt.Errorf("HTTP server[%s] : %w", s.Srv.Addr, err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context, cancel context.CancelFunc) {
	s.Log.Info("Shutting down HTTPServer...")
	_ = s.Srv.Shutdown(ctx)
	cancel()
	s.Log.Debug("HTTPServer shut down")
}
