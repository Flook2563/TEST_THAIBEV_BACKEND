package domain

import (
	"encoding/json"
)

type CreateUserProfileRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Profile    string `json:"profile"`
	Occupation string `json:"occupation"`
	BirthDay   string `json:"birth_day"`
	Sex        string `json:"sex"`
}

type CreateUserProfileResponse struct {
	UserID string `json:"user_id"`
}

type UserProfileRequest struct {
	UserID string `json:"	"`
}

type UserProfileResponse struct {
	UserID     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Profile    string `json:"profile"`
	Occupation string `json:"occupation"`
	BirthDay   string `json:"birth_day"`
	Sex        string `json:"sex"`
}

type AllUserProfileResponse struct {
	UserID     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Profile    string `json:"profile"`
	Occupation string `json:"occupation"`
	BirthDay   string `json:"birth_day"`
	Sex        string `json:"sex"`
	Createdate string `json:"create_date"`
}

func (r *CreateUserProfileRequest) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserProfileResponse) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserProfileRequest) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UserProfileResponse) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
func (r *AllUserProfileResponse) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
