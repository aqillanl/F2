package service

import (
	"errors"
	"finalproject2/model"
	"finalproject2/repository"
)

type SocialMediaService interface {
	FindAll(UserID int) ([]model.SocialMedia, error)
	CreateSocialMedia(UserID int, socialmedia model.SocialMediaInput) (model.SocialMedia, error)
	UpdateSocialMedia(id int, UserID int, socialmedia model.SocialMediaInput) (model.SocialMedia, error)
	DeleteSocialMedia(id, UserID int) error
}

type socialMediaService struct {
	repository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{socialMediaRepository}
}

// func (s *socialMediaService) FindAll(UserID int) ([]model.SocialMedia, error) {
// 	socialmedia, err := s.repository.FindAll(UserID)
// 	if err != nil {
// 		return socialmedia, err
// 	}
// 	return socialmedia, nil
// }

func (s *socialMediaService) FindAll(UserID int) ([]model.SocialMedia, error) {
	var socialMedias []model.SocialMedia
	err := s.repository.FindAll(&socialMedias)
	if err != nil {
		return socialMedias, err
	}

	if len(socialMedias) == 0 {
		return socialMedias, errors.New("no social media found")
	}

	var filteredSocialMedias []model.SocialMedia
	for _, socialMedia := range socialMedias {
		if socialMedia.UserID == UserID {
			filteredSocialMedias = append(filteredSocialMedias, socialMedia)
		}
	}

	return filteredSocialMedias, nil
}

func (s *socialMediaService) ShowSocialMedia(id int) (model.SocialMedia, error) {
	socialmedia, err := s.repository.FindById(id)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return socialmedia, nil
}

func (s *socialMediaService) CreateSocialMedia(Userid int, newSocialMedia model.SocialMediaInput) (model.SocialMedia, error) {

	var socialMedia = model.SocialMedia{
		Name:           newSocialMedia.Name,
		SocialMediaURL: newSocialMedia.SocialMediaURL,
		UserID:         Userid,
	}

	socialMediaResponse, err := s.repository.Create(socialMedia)
	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialMediaResponse, nil

}

func (s *socialMediaService) UpdateSocialMedia(id int, UserID int, socialmedia model.SocialMediaInput) (model.SocialMedia, error) {
	updatedSocialMedia, err := s.repository.FindById(id)
	if err != nil {
		return model.SocialMedia{}, err
	}

	if updatedSocialMedia.UserID != UserID {
		err = errors.New("you are not allowed to update this id")
		return model.SocialMedia{}, err
	}

	updatedSocialMedia.Name = socialmedia.Name
	updatedSocialMedia.SocialMediaURL = socialmedia.SocialMediaURL
	updatedSocialMedia, err = s.repository.Update(updatedSocialMedia)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return updatedSocialMedia, nil
}

func (s *socialMediaService) DeleteSocialMedia(id, UserID int) error {
	socialmedia, err := s.repository.FindById(id)
	if err != nil {
		return err
	}

	if socialmedia.UserID != UserID {
		err = errors.New("you are not allowed to delete this id")
		return err
	}

	err = s.repository.Delete(socialmedia)
	if err != nil {
		return err
	}
	return nil
}
