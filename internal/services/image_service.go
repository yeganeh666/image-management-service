package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"image-management-service/internal/models"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type ImageService interface {
	Upload(file *multipart.FileHeader, wg sync.WaitGroup) error
	List() ([]*models.Image, error)
	Get(id string) ([]byte, error)
}

type ImageServiceImpl struct {
	*Service
}

func NewImageService(service *Service) ImageService {
	return &ImageServiceImpl{
		Service: service,
	}
}

func (s ImageServiceImpl) Upload(file *multipart.FileHeader, wg sync.WaitGroup) error {
	defer wg.Done()

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		log.Errorf("failed to open file %q: %v", file.Filename, err)
		return err
	}
	defer src.Close()

	// Create the destination file
	filename := filepath.Base(file.Filename)
	ext := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(ext)]
	timestamp := time.Now().Format("20060102150405")
	dstName := fmt.Sprintf("%s-%s%s", name, timestamp, ext)
	path := filepath.Join(s.Config.Image.UploadPath, dstName)
	dst, err := os.Create(path)
	if err != nil {
		log.Errorf("failed to create file %q: %v", dstName, err)
		return err
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the destination file
	size, err := io.Copy(dst, src)
	if err != nil {
		log.Errorf("failed to copy file %q: %v", file.Filename, err)
		return err
	}

	// Save the image detail to datastore
	image := &models.Image{
		OriginalURL:   "",
		LocalName:     dstName,
		Path:          path,
		FileExtension: ext,
		FileSize:      size,
		DownloadDate:  time.Now().UTC(),
	}
	return s.ImageRepository.Save(image)
}

func (s ImageServiceImpl) List() ([]*models.Image, error) {
	// Read images directory
	//files, err := ioutil.ReadDir("./images")
	//if err != nil {
	//	log.WithError(err).Error("failed to read images directory")
	//	return nil, err
	//}
	//
	//// Convert files to Image struct
	//images := make([]*models.Image, 0)
	//for _, file := range files {
	//	// Extract image details from file name
	//	localName := file.Name()
	//	fileExt := filepath.Ext(localName)
	//	originalURL := "https://example.com/" + localName // Replace with actual URL
	//	fileSize := file.Size()
	//	downloadDate := file.ModTime().UTC()
	//
	//	// Add image to slice
	//	images = append(images, &models.Image{
	//		OriginalURL:   originalURL,
	//		LocalName:     localName,
	//		FileExtension: fileExt,
	//		FileSize:      fileSize,
	//		DownloadDate:  downloadDate,
	//	})
	//}
	//return images, nil
	return s.ImageRepository.List()
}

func (s ImageServiceImpl) Get(id string) ([]byte, error) {
	image, err := s.ImageRepository.Get(id)
	if err != nil {
		log.WithError(err).Error("failed to get image")
		return nil, err
	}

	data, err := ioutil.ReadFile(s.Config.Image.DownloadPath + image.LocalName)
	if err != nil {
		log.WithError(err).Error("failed to read images directory")
		return nil, err
	}

	return data, nil
}
