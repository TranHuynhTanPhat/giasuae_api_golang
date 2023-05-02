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

type ClassController interface {
	InsertClass(context *gin.Context)
	UpdateClass(context *gin.Context)
	DeleteClass(context *gin.Context)
	FindAllClass(context *gin.Context)
	FindByID(context *gin.Context)
}

type classController struct {
	ClassService services.ClassService
}

// DeleteClass implements ClassController
func (ctrl *classController) DeleteClass(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có lớp học được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err2 := ctrl.ClassService.DeleteClass(int(id))
	if err2 != nil {
		res := helper.BuildResponseError("Xóa lớp học thất bại", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FindAllClass implements ClassController
func (ctrl *classController) FindAllClass(context *gin.Context) {
	var classes []entities.Class = ctrl.ClassService.FindAllClass()
	res := helper.BuildResponse(true, "OK", classes)
	context.JSON(http.StatusOK, res)
}

// FindByID implements ClassController
func (ctrl *classController) FindByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không tìm thấy id", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var classes entities.Class = ctrl.ClassService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", classes)
	context.JSON(http.StatusOK, res)
}

// InsertClass implements ClassController
func (ctrl *classController) InsertClass(context *gin.Context) {
	var class entities.Class = entities.Class{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.BindJSON(&class)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errsub := ctrl.ClassService.InsertClass(&class)
	if errsub != nil {
		res := helper.BuildResponseError("Lớp học đã tồn tại", errsub.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", class)
	context.JSON(http.StatusOK, res)
}

// UpdateClass implements ClassController
func (ctrl *classController) UpdateClass(context *gin.Context) {
	var class entities.Class

	err := context.ShouldBind(&class)
	if err != nil {
		res := helper.BuildResponseError("Cập nhật môn học thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.ClassService.UpdateClass(&class)
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewClassITController(classSv services.ClassService) ClassController {
	return &classController{
		ClassService: classSv,
	}
}
