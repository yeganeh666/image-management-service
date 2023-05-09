package services

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"image-management-service/internal/models"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ImageService interface {
	Upload(file *multipart.FileHeader) error
	List() ([]*models.Image, error)
	Get(id string) ([]byte, error)
}

const localURL = "/api/images/"

type ImageServiceImpl struct {
	*Service
}

func NewImageService(service *Service) ImageService {
	return &ImageServiceImpl{
		Service: service,
	}
}

func (s ImageServiceImpl) Upload(file *multipart.FileHeader) error {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		s.log.WithError(err).Errorf("failed to open file %q", file.Filename)
		return err
	}
	defer src.Close()

	// Create the destination file
	filename := filepath.Base(file.Filename)
	ext := filepath.Ext(filename)
	imageID := uuid.New()
	path := filepath.Join(s.Config.Image.UploadPath, imageID.String())
	dst, err := os.Create(path)
	if err != nil {
		s.log.WithError(err).Errorf("failed to create file %q", filename)
		return err
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the destination file
	size, err := io.Copy(dst, src)
	if err != nil {
		s.log.WithError(err).Errorf("failed to copy file %q", file.Filename)
		return err
	}

	// Save the image detail to datastore
	image := &models.Image{
		OriginalURL:   "",
		OriginalName:  filename,
		LocalURL:      localURL + imageID.String(),
		Path:          path,
		FileExtension: ext,
		FileSize:      size,
		DownloadDate:  time.Now().UTC(),
	}
	image.ID = imageID
	return s.ImageRepository.Save(image)
}

func (s ImageServiceImpl) List() ([]*models.Image, error) {
	// Read images directory
	images, err := s.ImageRepository.List()
	if err != nil {
		log.WithError(err).Error("failed to read images directory")
		return nil, err
	}

	// Convert files to Image struct
	for _, image := range images {
		image.LocalURL = fmt.Sprint(
			s.Config.HTTP.Address, "/", s.Config.HTTP.Port, image.LocalURL) // Replace with actual URL
	}

	return images, nil
}

func (s ImageServiceImpl) Get(id string) ([]byte, error) {
	image, err := s.ImageRepository.Get(id)
	if err != nil {
		s.log.WithError(err).Error("failed to get image")
		return nil, err
	}

	data, err := ioutil.ReadFile("./" + image.Path)
	if err != nil {
		s.log.WithError(err).Error("failed to read images directory")
		return nil, err
	}

	return data, nil
}
