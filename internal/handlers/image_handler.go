package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

const (
	UploadPath = "./images/user-content/"
)

func (h *ImageHandlerImpl) HandleImagesUpload(c *gin.Context) {

	// Get list of uploaded files
	form, _ := c.MultipartForm()
	files := form.File["images"]

	// Create directory if not exist
	err := os.MkdirAll(UploadPath, 0755)
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
		go h.ImageService.Upload(file, wg)
	}

	// Wait for all the image uploads to finish before responding to the client
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "Upload successful"})
}

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
