package controller_admin

import (
	"context"
	"net/http"
	service_admin "server/application/service/admin"
	request_admin "server/presentation/request/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdController struct {
	AdApplicationService *service_admin.AdApplicationService
}

func NewAdController(AdApplicationService *service_admin.AdApplicationService) *AdController {
	return &AdController{AdApplicationService: AdApplicationService}
}

func (ctrl *AdController) FindAll(ctx *gin.Context) {
	var qs request_admin.AdFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.AdApplicationService.FindAll(service_admin.AdApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *AdController) FindByID(ctx *gin.Context) {
	adID, err := strconv.Atoi(ctx.Param("adID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.AdApplicationService.FindByID(service_admin.AdApplicationServiceFindByIDParams{
		Ctx:  context.Background(),
		AdID: adID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *AdController) Create(ctx *gin.Context) {
	var body request_admin.AdCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.AdApplicationService.Create(service_admin.AdApplicationServiceCreateParams{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *AdController) Update(ctx *gin.Context) {

	adID, err := strconv.Atoi(ctx.Param("adID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request_admin.AdUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.AdApplicationService.Update(service_admin.AdApplicationServiceUpdateParams{
		Ctx:  context.Background(),
		AdID: adID,
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *AdController) Delete(ctx *gin.Context) {

	adID, err := strconv.Atoi(ctx.Param("adID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.AdApplicationService.Delete(service_admin.AdApplicationServiceDeleteParams{
		Ctx:  context.Background(),
		AdID: adID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *AdController) UpdateIsActive(ctx *gin.Context) {

	adID, err := strconv.Atoi(ctx.Param("adID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request_admin.AdUpdateIsActiveRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.AdApplicationService.UpdateIsActive(service_admin.AdApplicationServiceUpdateIsActiveParams{
		Ctx:  context.Background(),
		AdID: adID,
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
