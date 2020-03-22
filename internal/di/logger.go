package di

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"

	"github.com/rmasclef/fizz_buzz_api/config"
)

func (container *Container) GetLogger() logger.WrappableLoggerInterface {
	if container.logger == nil {
		container.logger = logger.NewLogger(container.getLoggerHandler())
	}
	return container.logger
}

func (container *Container) getLoggerHandler() logger.HandlerInterface {
	// we will log on STDOUT with GELF format
	h := handler.Stream(os.Stdout, formatter.NewGelf())
	return container.getLoggerHandlerMiddleware().Decorate(h)
}

func (container *Container) getLoggerHandlerMiddleware() logger.Middlewares {
	return logger.MiddlewareStack(
		middleware.Placeholder(),
		middleware.Context(logger.Ctx("facility", config.AppName).Add("version", config.AppVersion)),
		middleware.MinLevelFilter(logger.LevelString(container.Cfg.LogLevel).Level()),
		middleware.Caller(3),
	)
}
