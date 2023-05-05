package handlers

import "image-management-service/internal/services"

type ImageHandlerImpl struct {
	ImageService      *services.ImageServiceImpl
	DownloaderService *services.DownloaderServiceImpl
}

func NewImageHandler(imageService *services.ImageServiceImpl, downloaderService *services.DownloaderServiceImpl) *ImageHandlerImpl {
	return &ImageHandlerImpl{
		ImageService:      imageService,
		DownloaderService: downloaderService,
	}
}
