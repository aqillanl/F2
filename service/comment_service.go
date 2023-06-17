package service

import (
	"errors"
	"finalproject2/model"
	"finalproject2/repository"
)

type CommentService interface {
	AddComment(id_user int, comment model.NewComment) (model.Comment, error)
	EditComment(id, id_user int, comment model.CommentUpdate) (model.Comment, error)
	DeleteComment(id, id_user int) error
	ViewComment(id_user int) ([]model.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) *commentService {
	return &commentService{
		commentRepository: commentRepo,
	}
}

func (c *commentService) AddComment(id_user int, comment model.NewComment) (model.Comment, error) {

	var newComment = model.Comment{
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID:  id_user,
	}

	commentResponse, err := c.commentRepository.AddComment(newComment)
	if err != nil {
		return model.Comment{}, err
	}

	return commentResponse, nil

	// cs := model.Comment{}
	// err := smapping.FillStruct(&cs, smapping.MapFields(&comment))
	// if err != nil {
	// 	return cs, err
	// }
	// result, err := c.commentRepository.AddComment(cs)
	// return result, err
}

func (c *commentService) EditComment(id, id_user int, comment model.CommentUpdate) (model.Comment, error) {
	updatedComment, err := c.commentRepository.FindById(id)
	if err != nil {
		return model.Comment{}, err
	}

	if updatedComment.UserID != id_user {
		err = errors.New("you are not allowed to edit this comment!")
		return model.Comment{}, err
	}

	updatedComment.Message = comment.Message

	result, err := c.commentRepository.EditComment(updatedComment)
	return result, err
}

func (c *commentService) DeleteComment(id, id_user int) error {
	deletedComment, err := c.commentRepository.FindById(id)
	if err != nil {
		return err
	}

	if deletedComment.UserID != id_user {
		err = errors.New("you are not allowed to delete this comment")
		return err
	}

	err = c.commentRepository.DeleteComment(deletedComment)
	if err != nil {
		return err
	}
	return nil
}

func (c *commentService) ViewComment(id int) ([]model.Comment, error) {
	comment, err := c.commentRepository.ViewComment(id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}
