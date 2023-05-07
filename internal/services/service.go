package services

import (
	log "github.com/sirupsen/logrus"
	"image-management-service/config"
	"image-management-service/internal/repositories"
)

type Service struct {
	log             *log.Logger
	Config          *config.Config
	ImageRepository repositories.ImageRepository
}

func NewService(log *log.Logger, config *config.Config, imageRepository repositories.ImageRepository) *Service {
	return &Service{
		log:             log,
		Config:          config,
		ImageRepository: imageRepository,
	}
}
