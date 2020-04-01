package di

import (
	"github.com/gol4ng/logger"

	"github.com/rmasclef/fizz_buzz_api/config"
	"github.com/rmasclef/fizz_buzz_api/internal/http/server"
)

type Container struct {
	Cfg *config.Config

	httpServer *server.Server

	logger logger.WrappableLoggerInterface
}

func NewContainer(cfg *config.Config) *Container {
	return &Container{Cfg: cfg}
}
