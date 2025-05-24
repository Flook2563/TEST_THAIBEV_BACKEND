package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"thaibev_backend/appconfig"
	"time"

	"github.com/labstack/echo/v4"
)

func ServerStart(cfg *appconfig.AppConfig) *echo.Echo {
	e := echo.New()

	// routes
	routes(e, cfg)

	go func() {
		endPoint := fmt.Sprintf(":%s", cfg.Server.Port)
		if err := e.Start(endPoint); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	return e
}

func ServerShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
