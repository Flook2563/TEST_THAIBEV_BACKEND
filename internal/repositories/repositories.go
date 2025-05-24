package repositories

import "gorm.io/gorm"

type Repo struct {
	TbTUserProfile TbTUserProfileRepo
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		TbTUserProfile: NewTbTUserRepo(db),
	}
}
