package routes

import (
	"finalproject2/handler"
	"finalproject2/middleware"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(router *gin.Engine, handler handler.PhotoHandler) {
	photos := router.Group("/photos")
	photos.Use(middleware.Authentication())
	photos.GET("/", handler.GetAllPhotoHandler)
	photos.POST("/", handler.CreatePhotoHandler)
	photos.PUT("/:id", handler.UpdatePhotoHandler)
	photos.DELETE("/:id", handler.DeletePhotoHandler)
}
