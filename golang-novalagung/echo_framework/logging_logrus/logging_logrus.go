package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func makeLogEntry(ctx echo.Context) *log.Entry {
	if ctx == nil {
		return log.WithFields(
			log.Fields{
				"at": time.Now().Format("2006-01-02 15:04:05"),
			},
		)
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": ctx.Request().Method,
		"uri":    ctx.Request().URL.String(),
		"ip":     ctx.Request().RemoteAddr,
	})
}

func main() {
	e := echo.New()
	e.Use(middlewareLogging)
	e.HTTPErrorHandler = errorHandler

	e.GET("/index", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, true)
	})

	lock := make(chan error)
	go func(lock chan error) {
		lock <- e.Start(":9000")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	makeLogEntry(nil).Warning("application started without ssl/tls enabled")

	err := <-lock
	if err != nil {
		makeLogEntry(nil).Panic("failed to start application")
	}
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		makeLogEntry(ctx).Info("incoming request")
		return next(ctx)
	}
}

func errorHandler(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(ctx).Error(report.Message)
	ctx.HTML(report.Code, report.Message.(string))
}
