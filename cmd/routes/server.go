package routes

import (
	"context"
	"fmt"
	"net/http"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Server forced to shutdown: ", err)
	}
}
