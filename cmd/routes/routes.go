package routes

import (
	"net/http"
	"thaibev_backend/appconfig"
	"thaibev_backend/database"
	"thaibev_backend/internal/handler"
	"thaibev_backend/internal/repositories"
	"thaibev_backend/internal/services"

	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo, cfg *appconfig.AppConfig) {
	db, dbErr := database.OpenPostgresqlDatabase(cfg.Database)
	if dbErr != nil {
		e.Logger.Fatal("Failed to connect to database: ", dbErr)
	}

	repo := repositories.NewRepository(db)
	Services := services.NewService(cfg, repo)
	handler := handler.NewHandler(Services, cfg)

	e.GET("/health", func(c echo.Context) error {
		response := map[string]string{
			"message": "service available",
		}
		return c.JSON(http.StatusOK, response)
	})

	v1 := e.Group("/api/v1")

	users := v1.Group("/users")
	users.POST("/profile", handler.CreateUserProfile)
	users.POST("/profile", handler.GetUserProfile)

}
