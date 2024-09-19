package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo Url is required"`
	Caption  string    `json:"caption" form:"caption" `
	Comments []Comment `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	UserId   uint
	User     *User
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	return
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	return
}
