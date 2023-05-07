package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"image-management-service/internal/handlers"
)

func Register(handler *handlers.ImageHandlerImpl) {
	route := gin.Default()
	route.Use(func(c *gin.Context) {
		c.Next()
	})

	r := route.Group("/api")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/images", handler.HandleImagesList)
	r.GET("/images/:id", handler.HandleDownloadImage)
	r.GET("/images/download", handler.HandleDownloadImages)
	r.POST("/images/upload", handler.HandleImagesUpload)

	// Start the server
	if err := route.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
