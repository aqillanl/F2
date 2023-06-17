package handler

import (
	"finalproject2/helpers"
	"finalproject2/model"
	"finalproject2/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUserHandler(ctx *gin.Context)
	LoginUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUserHandler(ctx *gin.Context) {

	var User model.User

	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	user, err := h.userService.RegisterUser(User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})

}

func (h *userHandler) LoginUserHandler(ctx *gin.Context) {

	var User model.User

	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	password := User.Password

	user, err := h.userService.LoginUser(User)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		log.Println(err.Error())
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		log.Println(err.Error())
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (h *userHandler) UpdateUserHandler(ctx *gin.Context) {

	var UserUpdate model.UserUpdateRequest

	err := ctx.ShouldBindJSON(&UserUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	id, _ := helpers.GetUserID(ctx)

	user, err := h.userService.UpdateUser(id, UserUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input a valid id",
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})

}

func (h *userHandler) DeleteUserHandler(ctx *gin.Context) {

	id, _ := helpers.GetUserID(ctx)

	err := h.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input a valid id",
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your account has been successfully deleted",
	})

}
