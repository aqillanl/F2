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

type SocialMediaHandler interface {
	CreateSocialMedia(ctx *gin.Context)
	GetSocialMedia(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}
type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) *socialMediaHandler {
	return &socialMediaHandler{socialMediaService}
}

func (h *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
	var input model.SocialMediaInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	id, _ := helpers.GetUserID(ctx)
	newSocialMedia, err := h.socialMediaService.CreateSocialMedia(id, input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatSocialMedia(newSocialMedia)
	ctx.JSON(http.StatusOK, formatter)
}

func (h *socialMediaHandler) GetSocialMedia(ctx *gin.Context) {
	id, _ := helpers.GetUserID(ctx)
	socialMedia, err := h.socialMediaService.FindAll(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatGetSocialMedia(socialMedia)
	ctx.JSON(http.StatusOK, formatter)
}

func (h *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	var inputData model.SocialMediaInput
	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	currentUserID, _ := helpers.GetUserID(ctx)

	updatedSocialMedia, err := h.socialMediaService.UpdateSocialMedia(id, currentUserID, inputData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatSocialMedia(updatedSocialMedia)
	ctx.JSON(http.StatusOK, formatter)
}

func (h *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
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

	err = h.socialMediaService.DeleteSocialMedia(id, currentUserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	formatter := model.FormatDeleteSocialMedia()
	ctx.JSON(http.StatusOK, formatter)
}
