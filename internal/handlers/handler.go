package handlers

import "image-management-service/internal/services"

type ImageHandlerImpl struct {
	ImageService      services.ImageService
	DownloaderService services.DownloaderService
}

func NewImageHandler(imageService services.ImageService, downloaderService services.DownloaderService) *ImageHandlerImpl {
	return &ImageHandlerImpl{
		ImageService:      imageService,
		DownloaderService: downloaderService,
	}
}
