package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"image-management-service/config"
	"image-management-service/internal/handlers"
	"image-management-service/internal/repositories"
	"image-management-service/internal/services"
)

const ServiceName = "ImageService"

func main() {
	log := logrus.New()
	conf, err := config.NewConfig(ServiceName)
	if err != nil {
		panic("failed to load service configuration")
	}

	repository, err := repositories.NewRepository(conf, log)
	if err != nil {
		panic("Failed to create repository")
	}

	imageRepository := repositories.NewImageRepositoryImpl(repository)
	service := services.NewService(imageRepository, log)
	handler := handlers.NewImageHandler(
		services.NewImageService(service),
		services.NewDownloaderService(service))

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	r.GET("/images", handler.HandleImagesList)
	r.GET("/images/:id", handler.HandleDownloadImage)
	r.GET("/images/download", handler.HandleImagesDownload)
	r.POST("/images/upload", handler.HandleImagesUpload)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
