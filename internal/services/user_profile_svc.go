package services

import (
	"context"
	"fmt"
	"thaibev_backend/internal/common"
	"thaibev_backend/internal/domain"
	"thaibev_backend/internal/repositories"
)

func (svc *service) CreateUserProfile(ctx context.Context, req *domain.CreateUserProfileRequest) (resp *domain.CreateUserProfileResponse, err error) {
	existing, err := svc.repo.TbTUserProfile.Search(ctx, repositories.TbTUserProfile{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	if len(existing) > 0 {
		return nil, fmt.Errorf("email already exists")
	}

	generatedUserID, err := svc.repo.TbTUserProfile.GenerateUserID(ctx)
	if err != nil {
		return nil, err
	}

	encryptedPhone, err := common.EncryptAES(req.Phone, svc.cfg.EncryptionKey)
	if err != nil {
		return nil, err
	}

	birthDay, err := common.ParseBirthDay(req.BirthDay)
	if err != nil {
		return nil, err
	}

	create, err := svc.repo.TbTUserProfile.Create(ctx, repositories.TbTUserProfile{
		Id:         generatedUserID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      encryptedPhone,
		Profile:    req.Profile,
		Occupation: req.Occupation,
		Sex:        req.Sex,
		BirthDay:   birthDay,
	})
	if err != nil {
		return nil, err
	}
	resp = &domain.CreateUserProfileResponse{
		UserID: create.Id,
	}

	return resp, nil
}

func (svc *service) GetUserProfile(ctx context.Context, req *domain.UserProfileRequest) (resp *domain.UserProfileResponse, err error) {

	user, err := svc.repo.TbTUserProfile.Search(ctx, repositories.TbTUserProfile{
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	User := user[0]

	decryptedPhone, err := common.DecryptAES(User.Phone, svc.cfg.EncryptionKey)
	if err != nil {
		return nil, err
	}

	resp = &domain.UserProfileResponse{
		UserID:     User.Id,
		FirstName:  User.FirstName,
		LastName:   User.LastName,
		Email:      User.Email,
		Phone:      decryptedPhone,
		Profile:    User.Profile,
		Occupation: User.Occupation,
		BirthDay:   User.BirthDay.String(),
		Sex:        User.Sex,
	}

	return resp, nil
}

func (svc *service) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	existing, err := svc.repo.TbTUserProfile.Search(ctx, repositories.TbTUserProfile{
		Email: email,
	})
	if err != nil {
		return false, err
	}
	return len(existing) > 0, nil
}

func (svc *service) DeleteUserProfile(ctx context.Context, userID string) error {
	if userID == "" {
		return fmt.Errorf("user_id is required")
	}
	return svc.repo.TbTUserProfile.Delete(ctx, repositories.TbTUserProfile{Id: userID})
}
