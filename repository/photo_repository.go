package repository

import (
	"finalproject2/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	ShowPhoto(id int) (model.Photo, error)
	ShowAllPhoto(photos *[]model.Photo) error
	CreatePhoto(photo model.Photo) (model.Photo, error)
	UpdatePhoto(photo model.Photo) (model.Photo, error)
	DeletePhoto(photo model.Photo) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) ShowPhoto(id int) (model.Photo, error) {
	var photo model.Photo
	err := r.db.Preload("User").Where("ID = ?", id).Take(&photo).Error
	if err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (r *photoRepository) ShowAllPhoto(photos *[]model.Photo) error {

	err := r.db.Preload("User").Find(&photos).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *photoRepository) CreatePhoto(photo model.Photo) (model.Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, err
}

func (r *photoRepository) UpdatePhoto(photo model.Photo) (model.Photo, error) {
	err := r.db.Debug().Save(&photo).Error
	if err != nil {
		return photo, err
	}
	r.db.Preload("User").Find(&photo)
	return photo, err
}

func (r *photoRepository) DeletePhoto(photo model.Photo) error {
	err := r.db.Debug().Delete(&photo).Error
	if err == nil {
		return err
	}
	return err
}
