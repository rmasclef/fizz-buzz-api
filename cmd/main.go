package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gol4ng/logger"

	"github.com/rmasclef/fizz_buzz_api/config"
	"github.com/rmasclef/fizz_buzz_api/internal/di"
)

func main() {
	fmt.Printf("Starting %s@%s on %s(%s/%s)\n", config.AppName, config.AppVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.New(ctx)

	// get DI
	container := di.NewContainer(cfg)
	log := container.GetLogger()

	// create chan that will receive stop reasons
	stopChan := make(chan interface{})

	// Start HTTP server
	httpSrv := container.GetHTTPServer(ctx)
	httpSrv.Start(stopChan)
	// intercept shutdown signals asynchronously
	handleShutdownSignal(stopChan)

	// wait for stop signal
	stopReason := <-stopChan
	logStopReason(stopReason, log)
	httpSrv.Stop(ctx, cancel)
}

func logStopReason(stopReason interface{}, log logger.LoggerInterface) {
	switch v := stopReason.(type) {
	case os.Signal:
		log.Info("received signal", logger.String("signal", v.String()))
	case error:
		log.Error("Fatal error", logger.Error("error", v))
	default:
		log.Warning(fmt.Sprintf("received unexpected stop value : %v", v))
	}
}

func handleShutdownSignal(stop chan<- interface{}) {
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		sig := <-signalChan
		stop <- sig
	}()
}
