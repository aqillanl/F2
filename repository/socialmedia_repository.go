package repository

import (
	"finalproject2/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	FindAll(socialmedia *[]model.SocialMedia) error
	FindById(id int) (model.SocialMedia, error)
	Create(socialmedia model.SocialMedia) (model.SocialMedia, error)
	Update(socialmedia model.SocialMedia) (model.SocialMedia, error)
	Delete(socialmedia model.SocialMedia) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r socialMediaRepository) FindAll(socialmedia *[]model.SocialMedia) error {
	err := r.db.Preload("User").Find(&socialmedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (r socialMediaRepository) FindById(id int) (model.SocialMedia, error) {
	var socialmedia model.SocialMedia
	err := r.db.Preload("User").Where("ID = ?", id).First(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}
	return socialmedia, nil
}

func (r socialMediaRepository) Create(socialmedia model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Create(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r socialMediaRepository) Update(socialmedia model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Debug().Save(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}
	return socialmedia, nil
}

func (r socialMediaRepository) Delete(socialmedia model.SocialMedia) error {
	err := r.db.Debug().Delete(&socialmedia).Error
	if err == nil {
		return err
	}
	return nil
}
