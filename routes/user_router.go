package routes

import (
	"finalproject2/handler"
	"finalproject2/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, handler handler.UserHandler) {
	user := router.Group("/user")
	user.POST("/register", handler.RegisterUserHandler)
	user.POST("/login", handler.LoginUserHandler)

	userData := router.Group("/user")
	userData.Use(middleware.Authentication())
	userData.PUT("/", handler.UpdateUserHandler)
	userData.DELETE("/", handler.DeleteUserHandler)
}
