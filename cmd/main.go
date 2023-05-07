package main

import (
	"github.com/sirupsen/logrus"
	"image-management-service/config"
	"image-management-service/internal/handlers"
	"image-management-service/internal/repositories"
	"image-management-service/internal/router"
	"image-management-service/internal/services"
)

const ServiceName = "ImageService"

func main() {
	log := logrus.New()
	conf, err := config.NewConfig(ServiceName)
	if err != nil {
		panic("failed to load service configuration")
	}

	repository, err := repositories.NewRepository(log, conf)
	if err != nil {
		panic("Failed to create repository")
	}

	service := services.NewService(log, conf,
		repositories.NewImageRepository(repository))

	handler := handlers.NewImageHandler(
		conf,
		services.NewImageService(service),
		services.NewDownloaderService(service))

	router.Routes(handler)
}
