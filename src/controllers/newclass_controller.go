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

type NewClassController interface {
	InsertNewClass(context *gin.Context)
	UpdateNewClass(context *gin.Context)
	DeleteNewClass(context *gin.Context)
	FindAllNewClass(context *gin.Context)
	FindByID(context *gin.Context)
	FilterNewClass(context *gin.Context)
	UpdateStatusNewClass(context *gin.Context)
}

type newClassController struct {
	NewClassService services.NewClassService
}

// UpdateStatusNewClass implements NewClassController
func (ctrl *newClassController) UpdateStatusNewClass(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Query("id"), 0, 0)
	status, _ := strconv.ParseUint(context.Query("status"), 0, 0)

	err2 := ctrl.NewClassService.UpdateStatusNewClass(int(status), int(id))
	if err2 != nil {
		res := helper.BuildResponseError("Cập nhật trạng thái thất bại", err2.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FilterNewClass implements NewClassController
func (ctrl *newClassController) FilterNewClass(context *gin.Context) {
	var newclasses []entities.NewclasssesSet
	subID, _ := strconv.ParseInt(context.Query("subID"), 0, 0)
	classID, _ := strconv.ParseInt(context.Query("classID"), 0, 0)
	cateID, _ := strconv.ParseInt(context.Query("cateID"), 0, 0)

	newclasses = ctrl.NewClassService.FilterNewClass(int(subID), int(classID), int(cateID))
	res := helper.BuildResponse(true, "OK", newclasses)
	context.JSON(http.StatusOK, res)
}

// DeleteNewClass implements NewClassController
func (ctrl *newClassController) DeleteNewClass(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có lớp mới được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err2 := ctrl.NewClassService.DeleteNewClass(int(id))
	if err2 != nil {
		res := helper.BuildResponseError("Xóa lớp mới thất bại", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FindAllNewClass implements NewClassController
func (ctrl *newClassController) FindAllNewClass(context *gin.Context) {
	page, _ := strconv.ParseInt(context.Query("page"), 0, 0)
	pagesize, _ := strconv.ParseInt(context.Query("pagesize"), 0, 0)

	var newclasses []entities.NewclasssesSet = ctrl.NewClassService.FindAllNewClass(int(page), int(pagesize))
	res := helper.BuildResponse(true, "OK", newclasses)
	context.JSON(http.StatusOK, res)
}

// FindByID implements NewClassController
func (ctrl *newClassController) FindByID(context *gin.Context) {
	var newclass entities.NewclassesDetail
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	newclass = ctrl.NewClassService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", newclass)
	context.JSON(http.StatusOK, res)
}

// InsertNewClass implements NewClassController
func (ctrl *newClassController) InsertNewClass(context *gin.Context) {
	var nc entities.NewClassesReq = entities.NewClassesReq{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.ShouldBind(&nc)
	if err != nil {
		res := helper.BuildResponseError("Không thể bind JSON", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	nc = ctrl.NewClassService.InsertNewClass(&nc)
	res := helper.BuildResponse(true, "OK", nc)
	context.JSON(http.StatusOK, res)
}

// UpdateNewClass implements NewClassController
func (ctrl *newClassController) UpdateNewClass(context *gin.Context) {
	var newclass entities.NewClassesReq
	if err := context.ShouldBind(&newclass); err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if err := ctrl.NewClassService.UpdateNewClass(&newclass); err != nil {
		res := helper.BuildResponseError("Cập nhật thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

func NewNewClassController(ncSvc services.NewClassService) NewClassController {
	return &newClassController{
		NewClassService: ncSvc,
	}
}
