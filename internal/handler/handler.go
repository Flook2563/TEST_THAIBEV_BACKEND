package handler

import (
	"thaibev_backend/appconfig"
	"thaibev_backend/internal/services"
)

type Handler interface {
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
