package controller_admin

import (
	"net/http"
	service_admin "server/application/service/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadCommentController struct {
	threadCommentApplicationService *service_admin.ThreadCommentApplicationService
}

func NewThreadCommentController(threadCommentApplicationService *service_admin.ThreadCommentApplicationService) *ThreadCommentController {
	return &ThreadCommentController{threadCommentApplicationService: threadCommentApplicationService}
}

func (ctrl *ThreadCommentController) Delete(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.threadCommentApplicationService.Delete(service_admin.ThreadCommentApplicationServiceFindByIDParams{
		Ctx:       ctx.Request.Context(),
		CommentID: commentID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
