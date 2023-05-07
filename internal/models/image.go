package models

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
	Path          string    `json:"path"`
}
