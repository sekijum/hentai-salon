package controller

import (
	"context"
	"net/http"
	"server/application/service"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagApplicationService *service.TagApplicationService
}

func NewTagController(tagApplicationService *service.TagApplicationService) *TagController {
	return &TagController{tagApplicationService: tagApplicationService}
}

func (ctrl *TagController) FindAllName(ginCtx *gin.Context) {
	tags, err := ctrl.tagApplicationService.FindAllName(service.TagApplicationServiceFindAllNameParams{
		Ctx: context.Background(),
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, tags)
}
