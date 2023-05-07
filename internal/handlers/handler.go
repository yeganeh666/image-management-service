package handlers

import (
	"image-management-service/config"
	"image-management-service/internal/services"
)

type ImageHandlerImpl struct {
	Config            *config.Config
	ImageService      services.ImageService
	DownloaderService services.DownloaderService
}

func NewImageHandler(config *config.Config, imageService services.ImageService, downloaderService services.DownloaderService) *ImageHandlerImpl {
	return &ImageHandlerImpl{
		Config:            config,
		ImageService:      imageService,
		DownloaderService: downloaderService,
	}
}
