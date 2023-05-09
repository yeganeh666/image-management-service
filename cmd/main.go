package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"image-management-service/config"
	"image-management-service/docs"
	"image-management-service/internal/handlers"
	"image-management-service/internal/repositories"
	"image-management-service/internal/router"
	"image-management-service/internal/services"
)

const ServiceName = "ImageService"

// @BasePath /api
// image management service godoc

func main() {
	log := logrus.New()
	conf, err := config.NewConfig(ServiceName)
	if err != nil {
		panic("failed to load service configuration")
	}

	initSwagger(conf)

	repository, err := repositories.NewRepository(log, conf)
	if err != nil {
		panic("Failed to create repository")
	}

	service := services.NewService(log, conf,
		repositories.NewImageRepository(repository))

	handler := handlers.NewImageHandler(
		log, conf,
		services.NewImageService(service),
		services.NewDownloaderService(service))

	router.Register(handler)
}
func initSwagger(conf *config.Config) {

	docs.SwaggerInfo.Title = "Image Management Service"
	docs.SwaggerInfo.Description = "Image Management Service : This is a simple upload/download service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("%v:%v", conf.HTTP.Address, conf.HTTP.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

}
