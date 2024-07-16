package controller

import (
	"context"
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

func (ctrl *StorageController) GeneratePresignedURLs(ginCtx *gin.Context) {
	var request struct {
		ObjectNames []string `json:"objectNames"`
	}

	if err := ginCtx.ShouldBindJSON(&request); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	urls, err := ctrl.storageApplicationService.GeneratePresignedURLs(service.StorageApplicationServiceGeneratePresignedURLs{
		Ctx:         context.Background(),
		ObjectNames: request.ObjectNames,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"urls": urls})
}
