package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"

	"github.com/gin-gonic/gin"
)

type ThreadCommentAttachmentController struct {
	threadCommentAttachmentApplicationService *service.ThreadCommentAttachmentApplicationService
}

func NewThreadCommentAttachmentController(threadCommentAttachmentApplicationService *service.ThreadCommentAttachmentApplicationService) *ThreadCommentAttachmentController {
	return &ThreadCommentAttachmentController{threadCommentAttachmentApplicationService: threadCommentAttachmentApplicationService}
}

func (ctrl *ThreadCommentAttachmentController) FindAll(ctx *gin.Context) {
	var qs request.ThreadCommentAttachmentFindAllRequest
	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadCommentAttachmentApplicationService.FindAll(service.ThreadCommentAttachmentApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
