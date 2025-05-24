package repositories

import "gorm.io/gorm"

type Repo struct {
	TbTUser TbTUserRepo
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		TbTUser: NewTbTUserRepo(db),
	}
}
