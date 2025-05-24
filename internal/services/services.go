package services

import (
	"thaibev_backend/appconfig"
	"thaibev_backend/internal/repositories"
)

type Service interface {
}

type service struct {
	cfg  *appconfig.AppConfig
	repo *repositories.Repo
}

func NewService(
	cfg *appconfig.AppConfig,
	repo *repositories.Repo,
) Service {
	return &service{
		cfg,
		repo,
	}
}
