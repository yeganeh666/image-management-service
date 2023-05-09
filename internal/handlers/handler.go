package handlers

import (
	"github.com/sirupsen/logrus"
	"image-management-service/config"
	"image-management-service/internal/services"
)

type ImageHandlerImpl struct {
	log               *logrus.Logger
	Config            *config.Config
	ImageService      services.ImageService
	DownloaderService services.DownloaderService
}

func NewImageHandler(log *logrus.Logger, config *config.Config, imageService services.ImageService, downloaderService services.DownloaderService) *ImageHandlerImpl {
	return &ImageHandlerImpl{
		log:               log,
		Config:            config,
		ImageService:      imageService,
		DownloaderService: downloaderService,
	}
}
