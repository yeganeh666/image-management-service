package main

import (
	log "github.com/sirupsen/logrus"
	"image-management-service/config"
	"image-management-service/internal/repositories"
)

const ServiceName = "ImageService"

func main() {
	conf := config.NewConfig(ServiceName)
	_, err := repositories.NewRepository(conf, log.New())
	if err != nil {
		log.WithError(err).Fatal("Failed to create repository")
	}
	return
}
