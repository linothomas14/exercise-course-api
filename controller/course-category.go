package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/service"
)

type CourseCategoryController interface {
	Create(context *gin.Context)
	FindAll(context *gin.Context)
	FindByID(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type courseCategoryController struct {
	courseCategoryService service.CourseCategoryService
}

func NewCourseCategoryController(courseCategoryService service.CourseCategoryService) CourseCategoryController {
	return &courseCategoryController{
		courseCategoryService: courseCategoryService,
	}
}

func (c *courseCategoryController) Create(ctx *gin.Context) {
	var param param.Category

	err := ctx.ShouldBind(&param)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(param)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	courseCategory := parseCourseCategory(param)

	res, err := c.courseCategoryService.CreateCourseCategory(courseCategory)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *courseCategoryController) FindAll(ctx *gin.Context) {

	res, err := c.courseCategoryService.FindAll()

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *courseCategoryController) FindByID(ctx *gin.Context) {

	id := ctx.Param("id")

	id64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	ctgID := uint32(id64)

	ctg, err := c.courseCategoryService.FindByID(ctgID)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		res := helper.BuildResponse("OK", ctg)
		ctx.JSON(http.StatusOK, res)
	}

}

func (c *courseCategoryController) Update(ctx *gin.Context) {

	var ctg model.CourseCategory

	err := ctx.ShouldBind(&ctg)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(ctg)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := ctx.Param("id")

	id64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	ctgID := uint32(id64)

	ctg.ID = ctgID

	res, err := c.courseCategoryService.Update(ctg)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *courseCategoryController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	id64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	ctgID := uint32(id64)

	err = c.courseCategoryService.Delete(ctgID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		res := helper.BuildResponse("Data has been deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	}

}

func parseCourseCategory(courseCategoryParam param.Category) model.CourseCategory {

	var courseCategory model.CourseCategory

	courseCategory.Name = courseCategoryParam.Name

	return courseCategory

}
