package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photos struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Your title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Your photo_url is required"`
	UserID   uint   `json:"user_id"`
	//User di skip validasinya
	User Users `gorm:"foreignkey:UserID" valid:"-"`
}

func (u *Photos) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
