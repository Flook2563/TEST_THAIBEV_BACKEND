package services

import (
	"context"
	"thaibev_backend/appconfig"
	"thaibev_backend/internal/domain"
	"thaibev_backend/internal/repositories"
)

type Service interface {
	CreateUserProfile(ctx context.Context, req *domain.CreateUserProfileRequest) (*domain.CreateUserProfileResponse, error)
	GetUserProfile(ctx context.Context, req *domain.UserProfileRequest) (*domain.UserProfileResponse, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	DeleteUserProfile(ctx context.Context, userID string) error
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
