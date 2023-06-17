package main

import (
	"finalproject2/config"
	"finalproject2/handler"
	"finalproject2/model"
	"finalproject2/repository"
	"finalproject2/routes"
	"finalproject2/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.Location)
	if err != nil {
		log.Println(err.Error())
		return
	}

	db.AutoMigrate(&model.User{}, &model.SocialMedia{}, &model.Photo{}, &model.Comment{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	routes.UserRouter(router, userHandler)

	socialRepository := repository.NewSocialMediaRepository(db)
	socialService := service.NewSocialMediaService(socialRepository)
	socialHandler := handler.NewSocialMediaHandler(socialService)

	routes.SocialMediaRouter(router, socialHandler)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService)

	routes.PhotoRouter(router, photoHandler)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	routes.CommentRouter(router, commentHandler)

	router.Run()
}
