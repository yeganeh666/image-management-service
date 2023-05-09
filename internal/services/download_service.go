package services

import (
	"github.com/google/uuid"
	"image-management-service/internal/models"
	"image-management-service/internal/utils/extention"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type DownloaderService interface {
	Download(url string) error
}

type DownloaderServiceImpl struct {
	*Service
}

func NewDownloaderService(service *Service) DownloaderService {
	return &DownloaderServiceImpl{
		Service: service,
	}
}

// Download an image from a given URL and save it to the specified directory
func (s DownloaderServiceImpl) Download(url string) error {
	downloadPath := s.Config.Image.DownloadPath
	// Create the directory if it doesn't exist
	err := os.MkdirAll(downloadPath, 0755)
	if err != nil {
		s.log.WithError(err).Error("Error creating directory")
		return err
	}

	// Send a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		s.log.WithError(err).Error("Error downloading image")
		return err
	}
	defer response.Body.Close()
	// Extract the file name from the URL
	fileName := filepath.Base(url)
	ext := extention.GetFileExtension(response)
	// Create a new file in the specified directory
	imageID := uuid.New()
	path := filepath.Join(downloadPath, imageID.String())
	file, err := os.Create(path)
	if err != nil {
		s.log.WithError(err).Error("Error creating file")
		return err
	}
	defer file.Close()

	// Copy the downloaded image to the file
	size, err := io.Copy(file, response.Body)
	if err != nil {
		s.log.WithError(err).Error("Error saving image")
		return err
	}

	image := &models.Image{
		OriginalURL:   url,
		OriginalName:  fileName,
		LocalURL:      localURL + imageID.String(),
		Path:          path,
		FileExtension: ext,
		FileSize:      size,
		DownloadDate:  time.Now(),
	}
	image.ID = imageID

	return s.ImageRepository.Save(image)
}
