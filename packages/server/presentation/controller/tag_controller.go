package controller

import (
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

func (ctrl *TagController) FindAll(c *gin.Context) {
	tags, err := ctrl.tagApplicationService.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "タグの取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}
