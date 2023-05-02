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

type PostController interface {
	InsertPost(context *gin.Context)
	UpdatePost(context *gin.Context)
	DeletePost(context *gin.Context)
	FindAllPost(context *gin.Context)
	FindByID(context *gin.Context)
	FilterPost(context *gin.Context)
}

type postController struct {
	PostService services.PostService
}

// FilterPost implements PostController
func (ctrl *postController) FilterPost(context *gin.Context) {
	types, _ := strconv.ParseInt(context.Query("type"), 0, 0)
	page, _ := strconv.ParseInt(context.Query("page"), 0, 0)
	pageSize, _ := strconv.ParseInt(context.Query("pagesize"), 0, 0)

	var posts []entities.Post = ctrl.PostService.FilterPost(int(types), int(page), int(pageSize))
	res := helper.BuildResponse(true, "OK", posts)
	context.JSON(http.StatusOK, res)
}

// DeletePost implements PostController
func (ctrl *postController) DeletePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không tìm thấy id", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	err1 := ctrl.PostService.DeletePost(int(id))
	if err1 != nil {
		res := helper.BuildResponseError("Xóa bài viết thất bại", err1.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", helper.EmptyObjec{})
	context.JSON(http.StatusOK, res)
}

// FindAllPost implements PostController
func (ctrl *postController) FindAllPost(context *gin.Context) {
	var postes []entities.Post = ctrl.PostService.FindAllPost()
	res := helper.BuildResponse(true, "OK", postes)
	context.JSON(http.StatusOK, res)
}

// FindByID implements PostController
func (ctrl *postController) FindByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không tìm thấy id", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var postes entities.Post = ctrl.PostService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", postes)
	context.JSON(http.StatusOK, res)
}

// InsertPost implements PostController
func (ctrl *postController) InsertPost(context *gin.Context) {
	var post entities.Post = entities.Post{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.BindJSON(&post)
	if err != nil {
		res := helper.BuildResponseError("Sai cú pháp", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	errsub := ctrl.PostService.InsertPost(&post)
	if errsub != nil {
		res := helper.BuildResponseError("Thêm bài viết thất bại", errsub.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", post)
	context.JSON(http.StatusOK, res)
}

// UpdatePost implements PostController
func (ctrl *postController) UpdatePost(context *gin.Context) {
	var post entities.Post

	err := context.ShouldBind(&post)
	if err != nil {
		res := helper.BuildResponseError("Cập nhật bài viết thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.PostService.UpdatePost(&post)
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewPostController(postSv services.PostService) PostController {
	return &postController{
		PostService: postSv,
	}
}
