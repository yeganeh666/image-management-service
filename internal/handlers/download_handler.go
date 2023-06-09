package handlers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"image-management-service/internal/utils/validation"
	"net/http"
	"os"
	"sync"
)

// HandleDownloadImages
// @Summary HandleDownloadImages
// @Description download images from links file
// @Tags Images
// @Produce json
// @Success 201
// @Router /images/download [get]
func (h *ImageHandlerImpl) HandleDownloadImages(c *gin.Context) {
	// Open the links.txt file
	file, err := os.Open(h.Config.Image.SourcePath)
	if err != nil {
		h.log.WithError(err).Error("Error opening file")
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
			go func() {
				defer wg.Done()
				err := h.DownloaderService.Download(url)
				if err != nil {
					h.log.WithError(err).Error("Failed to download image")
				}
			}()
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
