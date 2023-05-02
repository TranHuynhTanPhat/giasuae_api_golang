package controllers

import (
	"fmt"
	"giasuaeapi/src/date_picker"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/helper"
	"giasuaeapi/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SalaryinfoController interface {
	InsertSalaryinfo(context *gin.Context)
	UpdateSalaryinfo(context *gin.Context)
	DeleteSalaryinfo(context *gin.Context)
	FindAllSalaryinfo(context *gin.Context)
	FindByID(context *gin.Context)
	FindByType(context *gin.Context)
}

type salaryinfoController struct {
	SalaryinfoService services.SalaryinfoService
}

// DeleteSalaryinfo implements SalaryinfoController
func (ctrl *salaryinfoController) DeleteSalaryinfo(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err1 := ctrl.SalaryinfoService.DeleteSalaryinfo(int(id))
	if err1 != nil {
		res := helper.BuildResponseError("Xóa thông tin mức lương thất bại", err1.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FindAllSalaryinfo implements SalaryinfoController
func (ctrl *salaryinfoController) FindAllSalaryinfo(context *gin.Context) {
	var sals []entities.SalaryinfoView = ctrl.SalaryinfoService.FindAllSalaryinfo()
	res := helper.BuildResponse(true, "OK", sals)
	context.JSON(http.StatusOK, res)
}

// FindByID implements SalaryinfoController
func (ctrl *salaryinfoController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var sal entities.SalaryinfoDetail = ctrl.SalaryinfoService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", sal)
	context.JSON(http.StatusOK, res)
}

// FindByType implements SalaryinfoController
func (ctrl *salaryinfoController) FindByType(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("type"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var sal []entities.SalaryinfoView = ctrl.SalaryinfoService.FindByType(int(id))
	res := helper.BuildResponse(true, "OK", sal)
	context.JSON(http.StatusOK, res)
}

// InsertSalaryinfo implements SalaryinfoController
func (ctrl *salaryinfoController) InsertSalaryinfo(context *gin.Context) {
	var sal entities.Salaryinfo = entities.Salaryinfo{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.BindJSON(&sal)
	fmt.Println(sal)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errsal := ctrl.SalaryinfoService.InsertSalaryinfo(&sal)
	if errsal != nil {
		res := helper.BuildResponseError("Thêm thông tin mức lương thất bại", errsal.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", sal)
	context.JSON(http.StatusOK, res)
}

// UpdateSalaryinfo implements SalaryinfoController
func (ctrl *salaryinfoController) UpdateSalaryinfo(context *gin.Context) {
	var sub entities.Salaryinfo
	err := context.BindJSON(&sub)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errsub := ctrl.SalaryinfoService.UpdateSalaryinfo(&sub)
	if errsub != nil {
		res := helper.BuildResponseError("Cập nhật thất bại", errsub.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewSalaryinfoController(subSv services.SalaryinfoService) SalaryinfoController {
	return &salaryinfoController{
		SalaryinfoService: subSv,
	}
}
