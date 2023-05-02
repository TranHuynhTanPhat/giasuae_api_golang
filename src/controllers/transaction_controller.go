package controllers

import (
	"giasuaeapi/src/date_picker"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/helper"
	"giasuaeapi/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransController interface {
	InsertTrans(context *gin.Context)
	FindAllTrans(context *gin.Context)
	FindByIDAcc(context *gin.Context)
	FilterTrans(context *gin.Context)
	Statistics(context *gin.Context)
}

type transController struct {
	TransService services.TransService
}

// Statistics implements TransController
func (ctrl *transController) Statistics(context *gin.Context) {
	day := context.Query("day")
	month := context.Query("month")
	year := context.Query("year")
	var statistic = entities.Statistics{
		Created_at: date_picker.FormatDataNow(),
	}
	ctrl.TransService.Statistics(&statistic, month, year, day)
	res := helper.BuildResponse(true, "OK", statistic)
	context.JSON(http.StatusOK, res)
}

// FilterTrans implements TransController
func (ctrl *transController) FilterTrans(context *gin.Context) {
	key := context.Query("key")
	key = "%" + key + "%"
	var trans []entities.Transactionhistories = ctrl.TransService.FilterTrans(key)
	res := helper.BuildResponse(true, "OK", trans)
	context.JSON(http.StatusOK, res)
}

// FindAllTrans implements TransController
func (ctrl *transController) FindAllTrans(context *gin.Context) {
	var transs []entities.Transactionhistories = ctrl.TransService.FindAllTrans()
	res := helper.BuildResponse(true, "OK", transs)
	context.JSON(http.StatusOK, res)
}

// FindByIDAcc implements TransController
func (ctrl *transController) FindByIDAcc(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var trans entities.Transactionhistories = ctrl.TransService.FindByIDAcc(int(id))
	res := helper.BuildResponse(true, "OK", trans)
	context.JSON(http.StatusOK, res)
}

// InsertTrans implements TransController
func (ctrl *transController) InsertTrans(context *gin.Context) {
	var trans entities.TransactionhistoriesReq = entities.TransactionhistoriesReq{
		Created_at: date_picker.FormatDataNow(),
	}
	err := context.BindJSON(&trans)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errTrans := ctrl.TransService.InsertTrans(&trans)
	if errTrans != nil {
		res := helper.BuildResponseError("Thêm lịch giao dịch thất bại", errTrans.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", trans)
	context.JSON(http.StatusOK, res)
}

func NewTransController(svc services.TransService) TransController {
	return &transController{
		TransService: svc,
	}
}
