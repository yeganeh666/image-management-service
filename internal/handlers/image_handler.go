package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

// HandleImagesUpload
// @Summary HandleImagesUpload
// @Description upload images
// @Tags Images
// @Param file	formData file true "images"
//
//	@Accept	multipart/form-data
//
// @Produce json
// @Success 200
// @Router /images/upload [post]
func (h *ImageHandlerImpl) HandleImagesUpload(c *gin.Context) {

	// Get list of uploaded files
	form, _ := c.MultipartForm()
	files := form.File["images"]

	// Create directory if not exist
	err := os.MkdirAll(h.Config.Image.UploadPath, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create directory",
		})
		return
	}
	var wg sync.WaitGroup
	// Loop through the uploaded files and handle each one concurrently
	for _, file := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := h.ImageService.Upload(file)
			if err != nil {
				h.log.WithError(err).Error("Failed to upload image")
			}
		}()
	}

	// Wait for all the image uploads to finish before responding to the client
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "Upload successful"})
}

// HandleImagesList
// @Summary HandleImagesList
// @Description images list
// @Tags Images
// @Produce json
// @Success 200 {array} models.Image
// @Router /images [get]
func (h *ImageHandlerImpl) HandleImagesList(c *gin.Context) {
	images, err := h.ImageService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read image file",
		})
		return
	}
	c.JSON(http.StatusOK, images)
}

// HandleDownloadImage
// @Summary HandleDownloadImage
// @Description downoald an image from list
// @Tags Images
// @Param id path string true "image ID"
// @Produce octet-stream
// @Success 200
// @Router /images/{id} [get]
func (h *ImageHandlerImpl) HandleDownloadImage(c *gin.Context) {
	id := c.Param("id")
	// Read file from images directory
	data, err := h.ImageService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read image file",
		})
		return
	}

	// Set headers for file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+id)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	// Send file as response
	c.Data(http.StatusOK, "application/octet-stream", data)
}
