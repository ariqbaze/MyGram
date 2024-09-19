package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserId  uint
	PhotoId uint
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	return
}

func (u *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	return
}
