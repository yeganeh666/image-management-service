package services

type ImageServiceImpl struct {
	*Service
}

func NewImageService(service *Service) *ImageServiceImpl {
	return &ImageServiceImpl{
		Service: service,
	}
}
