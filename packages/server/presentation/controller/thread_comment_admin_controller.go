package controller

import (
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadCommentAdminController struct {
	threadCommentAdminApplicationService *service.ThreadCommentAdminApplicationService
}

func NewThreadCommentAdminController(threadCommentAdminApplicationService *service.ThreadCommentAdminApplicationService) *ThreadCommentAdminController {
	return &ThreadCommentAdminController{threadCommentAdminApplicationService: threadCommentAdminApplicationService}
}

func (ctrl *ThreadCommentAdminController) Update(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request.ThreadCommentAdminUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.threadCommentAdminApplicationService.Update(service.ThreadCommentAdminApplicationServiceFindByIDParams{
		Ctx:       ctx.Request.Context(),
		CommentID: commentID,
		Body:      body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
