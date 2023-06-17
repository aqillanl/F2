package service

import (
	"errors"
	"finalproject2/model"
	"finalproject2/repository"
)

type PhotoService interface {
	GetAllPhoto(UserID int) ([]model.Photo, error)
	CreatePhoto(UserID int, photo model.PhotoInput) (model.Photo, error)
	UpdatePhoto(id int, UserID int, photo model.PhotoInput) (model.Photo, error)
	DeletePhoto(id, UserID int) error
}

type photoService struct {
	repository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) *photoService {
	return &photoService{photoRepository}
}

func (s *photoService) GetAllPhoto(UserID int) ([]model.Photo, error) {
	var photos []model.Photo
	err := s.repository.ShowAllPhoto(&photos)
	if err != nil {
		return photos, err
	}
	if len(photos) == 0 {
		return photos, errors.New("no photo found")
	}
	var filteredPhotos []model.Photo
	for _, photo := range photos {
		if photo.UserID == UserID {
			filteredPhotos = append(filteredPhotos, photo)
		}
	}
	return filteredPhotos, nil
}

func (s *photoService) CreatePhoto(UserID int, newPhoto model.PhotoInput) (model.Photo, error) {
	photo := model.Photo{
		UserID:   UserID,
		Title:    newPhoto.Title,
		Caption:  newPhoto.Caption,
		PhotoUrl: newPhoto.PhotoUrl,
	}

	photoResponse, err := s.repository.CreatePhoto(photo)
	if err != nil {
		return model.Photo{}, err
	}
	return photoResponse, err
}

func (s *photoService) UpdatePhoto(id int, UserID int, photoUpdate model.PhotoInput) (model.Photo, error) {
	photo, err := s.repository.ShowPhoto(id)
	if err != nil {
		return model.Photo{}, err
	}
	if photo.UserID != UserID {
		err = errors.New("you are not allowed to update this id")
		return model.Photo{}, err
	}

	photo.Title = photoUpdate.Title
	photo.Caption = photoUpdate.Caption
	photo.PhotoUrl = photoUpdate.PhotoUrl

	updatePhoto, err := s.repository.UpdatePhoto(photo)
	if err != nil {
		return model.Photo{}, err
	}

	return updatePhoto, err
}

func (s *photoService) DeletePhoto(id int, UserID int) error {

	photo, err := s.repository.ShowPhoto(id)
	if err != nil {
		return err
	}

	if photo.UserID != UserID {
		err = errors.New("you are not allowed to delete this id")
		return err
	}

	err = s.repository.DeletePhoto(photo)
	if err != nil {
		return err
	}
	return nil
}
