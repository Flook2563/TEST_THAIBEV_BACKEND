package services

import (
	"context"
	"thaibev_backend/internal/domain"
	"thaibev_backend/internal/repositories"
)

func (svc *service) GetUserProfile(ctx context.Context, req *domain.UserProfileRequest) (resp *domain.UserProfileResponse, err error) {

	user, err := svc.repo.TbTUserProfile.Search(ctx, repositories.TbTUserProfile{
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	User := user[0]

	resp = &domain.UserProfileResponse{
		UserID: User.Id,
	}

	return resp, nil
}
