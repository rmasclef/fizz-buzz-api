package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)
type (
	fizzBuzzRequest struct {
		Str1  string `json:"str1" validate:"required"`
		Str2  string `json:"str2" validate:"required"`
		Int1  uint   `json:"int1" validate:"required"`
		Int2  uint   `json:"int2" validate:"required"`
		Limit uint   `json:"limit" validate:"required"`
	}
	fizzBuzzResponse []string
)

func (fr fizzBuzzRequest) String() string {
	return fmt.Sprintf("int1:%d_int2:%d_limit:%d_str1:%s_str2:%s", fr.Int1, fr.Int2, fr.Limit, fr.Str1, fr.Str2)
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// start HTTP server
	srv := getHTTPServer()
	srv.Logger.SetLevel(log.INFO)

	// Start server
	go func() {
		if err := srv.Start(":8080"); err != nil {
			srv.Logger.Error(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	srv.Logger.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}
}

func getHTTPServer() *echo.Echo {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	metricsCollector := &metricsCollector{}

	// Middlewares
	e.Use(middleware.Logger())  // enable logging
	e.Use(middleware.Recover()) // catch panics and recover from them (return an HTTP 500)

	e.POST("/fizz-buzz", fizzBuzzHandler(metricsCollector), requestContentTypeFilterer()) // only allow 'application/json' request content-type
	e.GET("/metrics", metricsHandler(metricsCollector))

	return e
}

// =====================================================================================================================
// ============================================================ MIDDLEWARES ============================================
// =====================================================================================================================

func requestContentTypeFilterer() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contentType := c.Request().Header.Get(echo.HeaderContentType)
			if contentType != echo.MIMEApplicationJSON && contentType != echo.MIMEApplicationJSONCharsetUTF8 {
				return c.JSON(
					http.StatusBadRequest,
					fmt.Sprintf(`This API only allows 'application/json' requests (provided: %s).`, contentType),
				)
			}
			return next(c)
		}
	}
}

// =====================================================================================================================
// ============================================================ HANDLERS ===============================================
// =====================================================================================================================

func metricsHandler(mc *metricsCollector) echo.HandlerFunc {
	return func(c echo.Context) error {
		if mc == nil || mc.RequestCounters == nil {
			return c.JSON(http.StatusOK, "no data collected yet")
		}
		sort.Sort(mc.RequestCounters)
		return c.JSON(http.StatusOK, mc.RequestCounters)
	}
}

func fizzBuzzHandler(mc *metricsCollector) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// bind HTTP Request (form-data or JSON) to fizzBuzzRequest
		req := new(fizzBuzzRequest)
		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		if err := ctx.Validate(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		// increment request counter metric
		if mc != nil {
			mc.IncRequestCounter(req.String())
		}

		// call fizzbuzz controller and transform the fizzbuzzResponse into an HTTP JSON expectedResponse
		return ctx.JSON(http.StatusOK, fizzBuzzController(req))
	}
}

func fizzBuzzController(fizzBuzzReq *fizzBuzzRequest) *fizzBuzzResponse {
	fizzBuzzResp := make(fizzBuzzResponse, fizzBuzzReq.Limit)

	for i := uint(1); i <= fizzBuzzReq.Limit; i++ {
		res := ""
		if i%fizzBuzzReq.Int1 == 0 {
			res += fizzBuzzReq.Str1
		}
		if i%fizzBuzzReq.Int2 == 0 {
			res += fizzBuzzReq.Str2
		}
		if res != "" {
			fizzBuzzResp[i-1] = res
			continue
		}
		fizzBuzzResp[i-1] = strconv.FormatUint(uint64(i), 10)
	}
	return &fizzBuzzResp
}
