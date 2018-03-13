package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	version string
	app     = kingpin.New("freedom-apiserver", "Freedom API server provides a REST API for managing Freedom proxy.").Version(version)
	port    = app.Flag("port", "Port to listen on.").Short('p').Default("8080").String()
	verbose = app.Flag("verbose", "Output debug log messages.").Short('v').Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	e := echo.New()
	e.Debug = *verbose

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Version string `json:"version"`
		}{
			version,
		})
	})

	go func() {
		if err := e.Start(":" + *port); err != nil {
			e.Logger.Info("shutting down server")
		}
	}()

	// Gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
