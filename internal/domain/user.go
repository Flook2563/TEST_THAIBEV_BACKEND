package domain

import "encoding/json"

type UserProfileRequest struct {
	UserID string `json:"user_id"`
}

type UserProfileResponse struct {
	UserID string `json:"user_id"`
}

func (r *UserProfileRequest) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserProfileResponse) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
