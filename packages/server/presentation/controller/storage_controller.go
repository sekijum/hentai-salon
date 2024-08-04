package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"

	"github.com/gin-gonic/gin"
)

type StorageController struct {
	storageApplicationService *service.StorageApplicationService
}

func NewStorageController(storageApplicationService *service.StorageApplicationService) *StorageController {
	return &StorageController{storageApplicationService: storageApplicationService}
}

func (ctrl *StorageController) GeneratePresignedURLs(ctx *gin.Context) {
	var body request.GeneratePresignedURLsRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urls, err := ctrl.storageApplicationService.GeneratePresignedURLs(service.StorageApplicationServiceGeneratePresignedURLs{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"urls": urls})
}
