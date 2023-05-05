package services

import (
	"fmt"
	"image-management-service/internal/repositories"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Downloader interface {
	Download(url string, dir string, wg *sync.WaitGroup) (*repositories.Image, error)
}

type DownloaderServiceImpl struct {
	*Service
}

func NewDownloaderService(service *Service) *DownloaderServiceImpl {
	return &DownloaderServiceImpl{
		Service: service,
	}
}

// Download an image from a given URL and save it to the specified directory
func (s DownloaderServiceImpl) Download(url string, dir string, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Create the directory if it doesn't exist
	err := os.MkdirAll(dir, 0755)
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
	file, err := os.Create(filepath.Join(dir, fileName))
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
	image := &repositories.Image{
		OriginalURL:   url,
		LocalName:     fileName,
		FileExtension: "",
		FileSize:      0,
		DownloadDate:  time.Now(),
	}
	fmt.Println("Downloaded:", fileName)
	return s.ImageRepository.SaveImage(image)
}
