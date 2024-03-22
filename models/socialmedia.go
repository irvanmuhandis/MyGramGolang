package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedias struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social_media_url is required"`
	UserId         uint   `json:"user_id"`
	User           Users  `gorm:"foreignkey:UserId" valid:"-"`
}

func (u *SocialMedias) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
