package controller_admin

import (
	"net/http"
	service_admin "server/application/service/admin"
	request_admin "server/presentation/request/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadCommentController struct {
	threadCommentApplicationService *service_admin.ThreadCommentApplicationService
}

func NewThreadCommentController(threadCommentApplicationService *service_admin.ThreadCommentApplicationService) *ThreadCommentController {
	return &ThreadCommentController{threadCommentApplicationService: threadCommentApplicationService}
}

func (ctrl *ThreadCommentController) Update(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request_admin.ThreadCommentUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.threadCommentApplicationService.Update(service_admin.ThreadCommentApplicationServiceFindByIDParams{
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
