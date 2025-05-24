package handler

import (
	"thaibev_backend/appconfig"
	"thaibev_backend/internal/services"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	CreateUserProfile(c echo.Context) error
	GetUserProfile(c echo.Context) error
}

type handler struct {
	services services.Service
	cfg      *appconfig.AppConfig
}

func NewHandler(
	services services.Service,
	config *appconfig.AppConfig,
) Handler {
	return &handler{
		services,
		config,
	}
}
