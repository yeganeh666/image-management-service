package repositories

import (
	"image-management-service/pkg/gormext"
	"time"
)

type Image struct {
	gormext.UniversalModel
	OriginalURL   string    `json:"original_url"`
	LocalName     string    `json:"local_name"`
	FileExtension string    `json:"file_extension"`
	FileSize      int64     `json:"file_size"`
	DownloadDate  time.Time `json:"download_date"`
}
type ImageRepository interface {
	SaveImage(*Image) (*Image, error)
}

type ImageRepositoryImpl struct {
	*Repository
}

func (r *ImageRepositoryImpl) SaveImage(img *Image) error {
	return r.DB.Save(img).Error
}

func NewImageRepositoryImpl(Repository *Repository) *ImageRepositoryImpl {
	return &ImageRepositoryImpl{
		Repository: Repository,
	}
}
