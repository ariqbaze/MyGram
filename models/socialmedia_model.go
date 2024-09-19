package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	UserId         uint
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social Media Url Url is required"`
	User           *User
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	return
}

func (u *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	return
}
