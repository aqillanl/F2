package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  int    `json:"user_id" form:"user_id"`
	PhotoID int    `json:"photo_id" form:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~your comment is require a message"`
	User    *User  `gorm:"foreignkey:UserID"`
	Photo   *Photo `gorm:"foreignkey:PhotoID"`
}

type CommentUpdate struct {
	// ID      uint   `json:"id" form:"id" valid:"required~Field ID is required"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~your comment is required"`
	UserID  int    `json:"user_id" valid:"required~Field ID is required"`
}

type NewComment struct {
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~your comment is required"`
	// UserID  uint   `json:"user_id" valid:"required~Field ID is required"`
	PhotoID int `json:"photo_id"`
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
