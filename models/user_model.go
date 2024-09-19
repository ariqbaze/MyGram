package models

import (
	"MyGram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your  username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Wrong email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,gte8~age must be 8 or older"`
	Photos       []Photo       `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	govalidator.CustomTypeTagMap.Set("gte8", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		if value, ok := i.(int); ok {
			return value >= 8
		}
		return false
	}))
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	govalidator.CustomTypeTagMap.Set("gte8", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		if value, ok := i.(int); ok {
			return value >= 8
		}
		return false
	}))
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}
