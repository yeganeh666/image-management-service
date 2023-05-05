package services

import (
	log "github.com/sirupsen/logrus"
	"image-management-service/internal/repositories"
)

type Service struct {
	log             *log.Logger
	ImageRepository *repositories.ImageRepositoryImpl
}

func NewService(imageRepository *repositories.ImageRepositoryImpl, log *log.Logger) *Service {
	return &Service{
		log:             log,
		ImageRepository: imageRepository,
	}
}
