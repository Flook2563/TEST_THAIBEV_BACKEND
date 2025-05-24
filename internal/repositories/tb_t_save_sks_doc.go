package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TbTUserRepo interface {
	Create(ctx context.Context, req TbTUser) (TbTUser, error)
	Search(ctx context.Context, filter TbTUser) ([]TbTUser, error)
	UpdateByFilter(ctx context.Context, filter TbTUser, update TbTUser) error
}

type tbUserRepo struct {
	db *gorm.DB
}

type TbTUser struct {
	Id         string    `gorm:"column:id;type:text;primaryKey" db:"id" json:"id"`
	FirstName  string    `gorm:"column:first_name;type:text" db:"first_name" json:"first_name"`
	LastName   string    `gorm:"column:last_name;type:text" db:"last_name" json:"last_name"`
	Email      string    `gorm:"column:email;type:text" db:"email" json:"email"`
	Phone      string    `gorm:"column:phone;type:text" db:"phone" json:"phone"`
	Profile    string    `gorm:"column:profile;type:text" db:"profile" json:"profile"`
	Occupation string    `gorm:"column:occupation;type:text" db:"occupation" json:"occupation"`
	BirthDay   time.Time `gorm:"column:birth_day;type:date" db:"birth_day" json:"birth_day"`
	Sex        string    `gorm:"column:sex;type:text" db:"sex" json:"sex"`
	CreateDate time.Time `gorm:"column:create_date;autoCreateTime" json:"create_date"`
	UpdateDate time.Time `gorm:"column:update_date;autoUpdateTime" json:"update_date"`
}

func (TbTUser) TableName() string {
	return "tb_t_user"
}

func NewTbTUserRepo(db *gorm.DB) TbTUserRepo {
	return &tbUserRepo{
		db: db,
	}
}

func (repo *tbUserRepo) Create(ctx context.Context, req TbTUser) (TbTUser, error) {
	query := repo.db.WithContext(ctx).Model(&TbTUser{})
	result := query.Create(&req)
	if result.Error != nil {
		return TbTUser{}, result.Error
	}
	return req, nil
}

func (repo *tbUserRepo) Search(ctx context.Context, filter TbTUser) ([]TbTUser, error) {
	query := repo.db.WithContext(ctx).Model(&TbTUser{}).Order("id DESC")

	if filter.Id != "" {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.FirstName != "" {
		query = query.Where("first_name = ?", filter.FirstName)
	}
	if filter.LastName != "" {
		query = query.Where("last_name = ?", filter.LastName)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	if filter.Phone != "" {
		query = query.Where("phone = ?", filter.Phone)
	}
	if filter.Profile != "" {
		query = query.Where("profile = ?", filter.Profile)
	}
	if filter.Occupation != "" {
		query = query.Where("occupation = ?", filter.Occupation)
	}
	if !filter.BirthDay.IsZero() {
		query = query.Where("birth_day = ?", filter.BirthDay)
	}
	if filter.Sex != "" {
		query = query.Where("sex = ?", filter.Sex)
	}

	resp := []TbTUser{}
	if err := query.Find(&resp).Error; err != nil {
		return resp, err
	}

	return resp, nil
}

func (repo *tbUserRepo) UpdateByFilter(
	ctx context.Context,
	filter TbTUser,
	update TbTUser,
) error {
	query := repo.db.WithContext(ctx).Model(&TbTUser{})

	if filter.Id != "" {
		query = query.Where("id = ?", filter.Id)
	}

	return query.Updates(update).Error
}
