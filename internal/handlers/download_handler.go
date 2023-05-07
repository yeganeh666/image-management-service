package handlers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"image-management-service/internal/utils/validation"
	"net/http"
	"os"
	"sync"
)

const (
	DirectoryName   = "images"
	sourceDirectory = "./data/links.txt"
)

func (h *ImageHandlerImpl) HandleImagesDownload(c *gin.Context) {
	// Open the links.txt file
	file, err := os.Open(sourceDirectory)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening file")
		return
	}
	defer file.Close()

	// Read the file line by line
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		if validation.IsValidImageURL(url) {
			wg.Add(1)
			go h.DownloaderService.Download(url, DirectoryName, &wg)
		}
	}
	if err := scanner.Err(); err != nil {
		c.String(http.StatusInternalServerError, "Error reading file")
		return
	}

	// Wait for all downloads to complete
	wg.Wait()
	c.JSON(http.StatusCreated, nil)

}
