package repositories

import "image-management-service/internal/models"

type ImageRepository interface {
	Save(*models.Image) error
	Get(id string) (*models.Image, error)
	List() ([]*models.Image, error)
}

type ImageRepositoryImpl struct {
	*Repository
}

func NewImageRepository(Repository *Repository) ImageRepository {
	return &ImageRepositoryImpl{
		Repository: Repository,
	}
}

func (r *ImageRepositoryImpl) Save(img *models.Image) error {
	return r.DB.Save(img).Error
}

func (r *ImageRepositoryImpl) Get(id string) (*models.Image, error) {
	image := &models.Image{}
	err := r.DB.Model(image).Where("id = ?", id).Find(&image).Error
	return image, err
}

func (r *ImageRepositoryImpl) List() ([]*models.Image, error) {
	images := make([]*models.Image, 0)
	err := r.DB.Model(&models.Image{}).Find(&images).Error
	return images, err
}
