package routes

import (
	"net/http"
	"thaibev_backend/appconfig"
	"thaibev_backend/database"

	"github.com/labstack/echo"
)

func routes(e *echo.Echo, cfg *appconfig.AppConfig) {
	db, dbErr := database.OpenPostgresqlDatabase(cfg.Database)
	if dbErr != nil {
		e.Logger.Fatal("Failed to connect to database: ", dbErr)
	}

	e.GET("/health", func(c echo.Context) error {
		response := map[string]string{
			"message": "service available",
		}
		return c.JSON(http.StatusOK, response)
	})

	v1 := e.Group("/api/v1")

	users := v1.Group("/users")

}
