package models

import (
	"finalassignment/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Your email have incorrect format,"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password need to have minimal length of 6 character"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,age8~Your age must be above 8"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
