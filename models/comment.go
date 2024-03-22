package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Your message is required"`
	UserID  uint   `json:"user_id"`
	User    Users  `gorm:"foreignkey:UserID" valid:"-"`
	PhotoID uint   `json:"photo_id"`
	Photo   Photos `gorm:"foreignkey:PhotoID" valid:"-"`
}

func (u *Comments) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
