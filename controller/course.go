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

type CourseController interface {
	Create(context *gin.Context)
	FindAll(context *gin.Context)
	FindByID(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type courseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &courseController{
		courseService: courseService,
	}
}

func (c *courseController) FindAll(ctx *gin.Context) {

	res, err := c.courseService.FindAll()

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if res == nil {
		response := helper.BuildResponse("OK", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)

		return
	}
	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *courseController) FindByID(ctx *gin.Context) {

	id := ctx.Param("id")

	id64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	courseID := uint32(id64)

	res, err := c.courseService.FindByID(courseID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", res)

	ctx.JSON(http.StatusOK, response)
}

func (c *courseController) Create(ctx *gin.Context) {

	var course model.Course

	err := ctx.ShouldBind(&course)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(course)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.courseService.CreateCourse(course)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusCreated, response)
}
func (c *courseController) Update(ctx *gin.Context) {

	var courseParam param.UpdateCourse
	var course model.Course

	id := ctx.Param("id")

	courseID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = ctx.ShouldBind(&courseParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(courseParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	courseParam.ID = uint32(courseID)

	course = parseCourseUpdate(courseParam)

	res, err := c.courseService.Update(course)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("Updated", res)
	ctx.JSON(http.StatusOK, response)
}
func (c *courseController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	courseId := uint32(u64)

	err = c.courseService.Delete(courseId)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("Course id "+id+" was deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}

func parseCourseUpdate(courseParam param.UpdateCourse) model.Course {

	var course model.Course

	course.ID = courseParam.ID
	course.Title = courseParam.Title
	course.CourseCategoryId = courseParam.CourseCategoryId

	return course

}
