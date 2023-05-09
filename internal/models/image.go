package models

import (
	"image-management-service/pkg/gormext"
	"time"
)

type Image struct {
	gormext.UniversalModel
	OriginalURL   string    `json:"original_url"`
	OriginalName  string    `json:"original_name"`
	LocalURL      string    `json:"local_url"`
	FileExtension string    `json:"file_extension"`
	FileSize      int64     `json:"file_size"`
	DownloadDate  time.Time `json:"download_date"`
	Path          string    `json:"path"`
}
