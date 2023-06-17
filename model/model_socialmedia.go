package model

import (
	"time"
)

type SocialMedia struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null;size:150" json:"name"`
	SocialMediaURL string    `gorm:"not null;size:150" json:"social_media_url"`
	UserID         int       `gorm:"index" json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `gorm:"foreignKey:UserID" json:"user"`
}

type SocialMediaInput struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}

type GetSocialMediaInput struct {
	ID int `uri:"id" binding:"required"`
}

// func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(s)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }
