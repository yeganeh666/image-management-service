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
	conf := config.NewConfig(ServiceName)
	log := logrus.New()
	repository, err := repositories.NewRepository(conf, log)
	if err != nil {
		log.WithError(err).Fatal("Failed to create repository")
	}

	service := services.NewService(repositories.NewImageRepositoryImpl(repository), log)
	handler := handlers.NewImageHandler(
		services.NewImageService(service),
		services.NewDownloaderService(service))

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Next()
	})
	r.GET("/images/download", handler.HandleImagesDownload)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
