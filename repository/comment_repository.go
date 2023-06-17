package repository

import (
	"finalproject2/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	AddComment(comment model.Comment) (model.Comment, error)
	EditComment(comment model.Comment) (model.Comment, error)
	DeleteComment(comment model.Comment) error
	ViewComment(id int) ([]model.Comment, error)
	FindById(id int) (model.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) AddComment(comment model.Comment) (model.Comment, error) {
	err := c.db.Create(&comment).Error
	return comment, err
}

func (c *commentRepository) EditComment(comment model.Comment) (model.Comment, error) {
	err := c.db.Save(&comment).Error
	return comment, err
}

func (c *commentRepository) DeleteComment(comment model.Comment) error {
	err := c.db.Debug().Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) ViewComment(id int) ([]model.Comment, error) {
	var comments []model.Comment
	err := c.db.Debug().Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func (c *commentRepository) FindById(id int) (model.Comment, error) {
	var comment model.Comment
	err := c.db.Preload("User").Where("ID = ?", id).First(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}
