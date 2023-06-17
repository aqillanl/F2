package routes

import (
	"finalproject2/handler"
	"finalproject2/middleware"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(router *gin.Engine, handler handler.SocialMediaHandler) {
	socialMedia := router.Group("/socialmedias")
	socialMedia.Use(middleware.Authentication())
	socialMedia.POST("/", handler.CreateSocialMedia)
	socialMedia.GET("/", handler.GetSocialMedia)
	socialMedia.PUT("/:id", handler.UpdateSocialMedia)
	socialMedia.DELETE("/:id", handler.DeleteSocialMedia)
}
