package model

import (
	"time"
)

type PhotoFormatter struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
}

func FormatPhoto(post Photo) PhotoFormatter {
	formatter := PhotoFormatter{
		ID:        post.ID,
		Title:     post.Title,
		Caption:   post.Caption,
		PhotoUrl:  post.PhotoUrl,
		CreatedAt: post.CreatedAt,
		UserID:    post.UserID,
	}

	return formatter
}

type GetPhotoFormatter struct {
	Photo []struct {
		ID        uint      `json:"id"`
		Title     string    `json:"title"`
		Caption   string    `json:"caption"`
		PhotoUrl  string    `json:"photo_url"`
		CreatedAt time.Time `json:"created_at"`
		UserID    int       `json:"user_id"`
		UpdatedAt time.Time `json:"updated_at"`
		User      struct {
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"photo"`
}

func FormatGetPhoto(post []Photo) GetPhotoFormatter {
	var formatter GetPhotoFormatter
	for _, value := range post {
		formatter.Photo = append(formatter.Photo, struct {
			ID        uint      `json:"id"`
			Title     string    `json:"title"`
			Caption   string    `json:"caption"`
			PhotoUrl  string    `json:"photo_url"`
			CreatedAt time.Time `json:"created_at"`
			UserID    int       `json:"user_id"`
			UpdatedAt time.Time `json:"updated_at"`
			User      struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"user"`
		}{
			ID:        value.ID,
			Title:     value.Title,
			Caption:   value.Caption,
			PhotoUrl:  value.PhotoUrl,
			CreatedAt: value.CreatedAt,
			UserID:    value.UserID,
			UpdatedAt: value.UpdatedAt,
			User: struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				Email:    value.User.Email,
				Username: value.User.Username,
			},
		})
	}

	return formatter
}

type UpdatePhotoFormatter struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUpdatePhoto(post Photo) UpdatePhotoFormatter {
	formatter := UpdatePhotoFormatter{
		ID:        post.ID,
		Title:     post.Title,
		Caption:   post.Caption,
		PhotoUrl:  post.PhotoUrl,
		UserID:    post.UserID,
		UpdatedAt: post.UpdatedAt,
	}
	return formatter
}

type DeletePhotoFormatter struct {
	Message string `json:"message"`
}

func FormatDeletePhoto() DeleteFormatter {
	formatter := DeleteFormatter{
		Message: "Your photo has been successfully deleted",
	}

	return formatter
}
