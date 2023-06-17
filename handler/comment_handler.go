package handler

import (
	"finalproject2/helpers"
	"finalproject2/model"
	"finalproject2/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	AddComment(ctx *gin.Context)
	EditComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	ViewComment(ctx *gin.Context)
}

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentSer service.CommentService) *commentHandler {
	return &commentHandler{
		commentService: commentSer,
	}
}

func (c *commentHandler) AddComment(ctx *gin.Context) {
	var newcomment model.NewComment
	errnew := ctx.ShouldBind(&newcomment)
	if errnew != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errnew.Error(),
		})
		log.Println(errnew.Error())
		return
	}
	id, _ := helpers.GetUserID(ctx)

	comment, err := c.commentService.AddComment(id, newcomment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errnew.Error(),
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func (c *commentHandler) EditComment(ctx *gin.Context) {
	var editComment model.CommentUpdate

	err := ctx.ShouldBindJSON(&editComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	id_user, _ := helpers.GetUserID(ctx)
	comment, err := c.commentService.EditComment(id, id_user, editComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input a valid id",
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})
}

func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	id_user, _ := helpers.GetUserID(ctx)

	err := c.commentService.DeleteComment(id, id_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input a valid id",
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}

func (c *commentHandler) ViewComment(ctx *gin.Context) {
	id, _ := helpers.GetUserID(ctx)
	comment, err := c.commentService.ViewComment(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatGetComment(comment)
	ctx.JSON(http.StatusOK, formatter)
}
