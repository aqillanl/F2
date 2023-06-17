package routes

import (
	"finalproject2/handler"
	"finalproject2/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.Engine, handler handler.CommentHandler) {
	comment := router.Group("/comment")
	comment.Use(middleware.Authentication())
	comment.GET("/", handler.ViewComment)
	comment.POST("/", handler.AddComment)
	comment.PUT("/:id", handler.EditComment)
	comment.DELETE("/:id", handler.DeleteComment)
}
