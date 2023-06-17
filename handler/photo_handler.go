package handler

import (
	"finalproject2/helpers"
	"finalproject2/model"
	"finalproject2/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	GetAllPhotoHandler(ctx *gin.Context)
	CreatePhotoHandler(ctx *gin.Context)
	UpdatePhotoHandler(ctx *gin.Context)
	DeletePhotoHandler(ctx *gin.Context)
}

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) *photoHandler {
	return &photoHandler{photoService}
}

func (h *photoHandler) GetAllPhotoHandler(ctx *gin.Context) {
	id, _ := helpers.GetUserID(ctx)
	photo, err := h.photoService.GetAllPhoto(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := model.FormatGetPhoto(photo)
	ctx.JSON(http.StatusOK, formatter)
}

func (h *photoHandler) CreatePhotoHandler(ctx *gin.Context) {
	var input model.PhotoInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	id, _ := helpers.GetUserID(ctx)

	photo, err := h.photoService.CreatePhoto(id, input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatPhoto(photo)
	ctx.JSON(http.StatusOK, formatter)
}

func (h *photoHandler) UpdatePhotoHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	var inputData model.PhotoInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	currentUserID, _ := helpers.GetUserID(ctx)

	updatedPhoto, err := h.photoService.UpdatePhoto(id, currentUserID, inputData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatPhoto(updatedPhoto)
	ctx.JSON(http.StatusOK, formatter)
}

func (h *photoHandler) DeletePhotoHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	currentUserID, _ := helpers.GetUserID(ctx)
	fmt.Println(currentUserID)

	err = h.photoService.DeletePhoto(id, currentUserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatDeletePhoto()
	ctx.JSON(http.StatusOK, formatter)
}
