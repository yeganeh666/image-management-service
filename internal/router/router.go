package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"image-management-service/internal/handlers"
)

func Routes(handler *handlers.ImageHandlerImpl) {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	r.GET("/images", handler.HandleImagesList)
	r.GET("/images/:id", handler.HandleDownloadImage)
	r.GET("/images/download", handler.HandleDownloadImages)
	r.POST("/images/upload", handler.HandleImagesUpload)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
