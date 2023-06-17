package model

import (
	"time"
)

type GetCommentFormatter struct {
	Comment []struct {
		ID        uint      `json:"id"`
		Message   string    `json:"message"`
		PhotoID   int       `json:"caption"`
		UserID    int       `json:"user_id"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
		User      struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"user"`
		Photo struct {
			ID        uint   `json:"id"`
			Title     string `json:"title"`
			Caption   string `json:"caption"`
			Photo_Url string `json:"photo_url"`
			User_id   int    `json:"user_id"`
		} `json:"photo"`
	} `json:"comment"`
}

func FormatGetComment(post []Comment) GetCommentFormatter {
	var formatter GetCommentFormatter
	for _, value := range post {
		formatter.Comment = append(formatter.Comment, struct {
			ID        uint      `json:"id"`
			Message   string    `json:"message"`
			PhotoID   int       `json:"caption"`
			UserID    int       `json:"user_id"`
			UpdatedAt time.Time `json:"updated_at"`
			CreatedAt time.Time `json:"created_at"`
			User      struct {
				ID       uint   `json:"id"`
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"user"`
			Photo struct {
				ID        uint   `json:"id"`
				Title     string `json:"title"`
				Caption   string `json:"caption"`
				Photo_Url string `json:"photo_url"`
				User_id   int    `json:"user_id"`
			} `json:"photo"`
		}{
			ID:        value.ID,
			Message:   value.Message,
			PhotoID:   value.PhotoID,
			UserID:    value.UserID,
			UpdatedAt: value.UpdatedAt,
			CreatedAt: value.CreatedAt,
			User: struct {
				ID       uint   `json:"id"`
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				ID:       value.User.ID,
				Email:    value.User.Email,
				Username: value.User.Username,
			},
			Photo: struct {
				ID        uint   `json:"id"`
				Title     string `json:"title"`
				Caption   string `json:"caption"`
				Photo_Url string `json:"photo_url"`
				User_id   int    `json:"user_id"`
			}{
				ID:        value.Photo.ID,
				Title:     value.Photo.Title,
				Caption:   value.Photo.Caption,
				Photo_Url: value.Photo.PhotoUrl,
				User_id:   value.Photo.UserID,
			},
		})
	}

	return formatter
}

// type UpdatePhotoFormatter struct {
// 	ID        uint      `json:"id"`
// 	Title     string    `json:"title"`
// 	Caption   string    `json:"caption"`
// 	PhotoUrl  string    `json:"photo_url"`
// 	UserID    int       `json:"user_id"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

// func FormatUpdatePhoto(post Photo) UpdatePhotoFormatter {
// 	formatter := UpdatePhotoFormatter{
// 		ID:        post.ID,
// 		Title:     post.Title,
// 		Caption:   post.Caption,
// 		PhotoUrl:  post.PhotoUrl,
// 		UserID:    post.UserID,
// 		UpdatedAt: post.UpdatedAt,
// 	}
// 	return formatter
// }

// type DeletePhotoFormatter struct {
// 	Message string `json:"message"`
// }

// func FormatDeletePhoto() DeleteFormatter {
// 	formatter := DeleteFormatter{
// 		Message: "Your photo has been successfully deleted",
// 	}

// 	return formatter
// }
