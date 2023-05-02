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

type CategoryController interface {
	InsertCategory(context *gin.Context)
	UpdateCategory(context *gin.Context)
	DeleteCategory(context *gin.Context)
	FindAllCategory(context *gin.Context)
	FindByID(context *gin.Context)
	FilterCategorry(context *gin.Context)
}

type categoryController struct {
	CategoryService services.CategoryService
}

// FilterCategorry implements CategoryController
func (ctrl *categoryController) FilterCategorry(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("type"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không tìm thấy id", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var category []entities.Category = ctrl.CategoryService.FilterCategory(int(id))
	res := helper.BuildResponse(true, "OK", category)
	context.JSON(http.StatusOK, res)
}

// DeleteCategory implements CategoryController
func (ctrl *categoryController) DeleteCategory(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có danh mục được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err2 := ctrl.CategoryService.DeleteCategory(int(id))
	if err2 != nil {
		res := helper.BuildResponseError("Xóa danh mục thất bại", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FindAllCategory implements CategoryController
func (ctrl *categoryController) FindAllCategory(context *gin.Context) {
	var categories []entities.Category = ctrl.CategoryService.FindAllCategory()
	res := helper.BuildResponse(true, "OK", categories)
	context.JSON(http.StatusOK, res)
}

// FindByID implements CategoryController
func (ctrl *categoryController) FindByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không tìm thấy id", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var category entities.Category = ctrl.CategoryService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", category)
	context.JSON(http.StatusOK, res)
}

// InsertCategory implements CategoryController
func (ctrl *categoryController) InsertCategory(context *gin.Context) {
	var category entities.Category = entities.Category{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.BindJSON(&category)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errsub := ctrl.CategoryService.InsertCategory(&category)
	if errsub != nil {
		res := helper.BuildResponseError("Danh mục đã tồn tại", errsub.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", category)
	context.JSON(http.StatusOK, res)
}

// UpdateCategory implements CategoryController
func (ctrl *categoryController) UpdateCategory(context *gin.Context) {
	var category entities.Category
	err := context.ShouldBind(&category)
	if err != nil {
		res := helper.BuildResponseError("Cập nhật danh mục thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.CategoryService.UpdateCategory(&category)
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewCategoryController(CategorySv services.CategoryService) CategoryController {
	return &categoryController{
		CategoryService: CategorySv,
	}
}
