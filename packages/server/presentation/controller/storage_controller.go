package controller

import (
	"net/http"
	"server/application/service"

	"github.com/gin-gonic/gin"
)

type StorageController struct {
	storageApplicationService *service.StorageApplicationService
}

func NewStorageController(storageApplicationService *service.StorageApplicationService) *StorageController {
	return &StorageController{storageApplicationService: storageApplicationService}
}

func (ctrl *StorageController) GeneratePresignedURLs(c *gin.Context) {
	var request struct {
		ObjectNames []string `json:"objectNames"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	urls, err := ctrl.storageApplicationService.GeneratePresignedURLs(request.ObjectNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
