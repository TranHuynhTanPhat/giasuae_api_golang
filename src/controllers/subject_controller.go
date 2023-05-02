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

type SubjectController interface {
	InsertSubject(context *gin.Context)
	UpdateSubject(context *gin.Context)
	DeleteSubject(context *gin.Context)
	FindAllSubject(context *gin.Context)
	FindByID(context *gin.Context)
}

type subjectController struct {
	SubjectService services.SubjectService
}

// DeleteSubject implements SubjectController
func (ctrl *subjectController) DeleteSubject(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có môn học được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err2 := ctrl.SubjectService.DeleteSubject(int(id))
	if err2 != nil {
		res := helper.BuildResponseError("Xóa môn học thất bại", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FindAllSubject implements SubjectController
func (ctrl *subjectController) FindAllSubject(context *gin.Context) {
	var subjects []entities.Subject = ctrl.SubjectService.FindAllSubject()
	res := helper.BuildResponse(true, "OK", subjects)
	context.JSON(http.StatusOK, res)
}

// FindByID implements SubjectController
func (ctrl *subjectController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var sub entities.Subject = ctrl.SubjectService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", sub)
	context.JSON(http.StatusOK, res)
}

// InsertSubject implements SubjectController
func (ctrl *subjectController) InsertSubject(context *gin.Context) {
	var sub entities.Subject = entities.Subject{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.BindJSON(&sub)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errsub := ctrl.SubjectService.InsertSubject(&sub)
	if errsub != nil {
		res := helper.BuildResponseError("Môn học đã tồn tại", errsub.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", sub)
	context.JSON(http.StatusOK, res)
}

// UpdateSubject implements SubjectController
func (ctrl *subjectController) UpdateSubject(context *gin.Context) {
	var sub entities.Subject

	err := context.ShouldBind(&sub)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.SubjectService.UpdateSubject(&sub)
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewSubjectController(subSv services.SubjectService) SubjectController {
	return &subjectController{
		SubjectService: subSv,
	}
}
