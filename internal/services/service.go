package services

import (
	log "github.com/sirupsen/logrus"
	"image-management-service/internal/repositories"
)

type Service struct {
	log             *log.Logger
	ImageRepository repositories.ImageRepository
}

func NewService(imageRepository repositories.ImageRepository, log *log.Logger) *Service {
	return &Service{
		log:             log,
		ImageRepository: imageRepository,
	}
}
