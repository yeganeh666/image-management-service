package services

import (
	"fmt"
	"image-management-service/internal/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type DownloaderService interface {
	Download(url string, wg *sync.WaitGroup) error
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
func (s DownloaderServiceImpl) Download(url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	downloadPath := s.Config.Image.DownloadPath
	// Create the directory if it doesn't exist
	err := os.MkdirAll(downloadPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	// Send a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return err
	}
	defer response.Body.Close()

	// Extract the file name from the URL
	fileName := filepath.Base(url)

	// Create a new file in the specified directory
	path := filepath.Join(downloadPath, fileName)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Copy the downloaded image to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return err
	}
	image := &models.Image{
		OriginalURL:   url,
		LocalName:     fileName,
		Path:          path,
		FileExtension: "",
		FileSize:      0,
		DownloadDate:  time.Now(),
	}
	fmt.Println("Downloaded:", fileName)
	return s.ImageRepository.Save(image)
}
