package controller_admin

import (
	"context"
	"net/http"
	service_admin "server/application/service/admin"
	request_admin "server/presentation/request/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagApplicationService *service_admin.TagApplicationService
}

func NewTagController(tagApplicationService *service_admin.TagApplicationService) *TagController {
	return &TagController{tagApplicationService: tagApplicationService}
}

func (ctrl *TagController) FindAll(ctx *gin.Context) {
	var qs request_admin.TagFindAllRequest
	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.tagApplicationService.FindAll(service_admin.TagApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *TagController) Delete(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tagID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.tagApplicationService.Delete(service_admin.TagApplicationServiceDeleteParams{
		Ctx:   context.Background(),
		TagID: tagID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
