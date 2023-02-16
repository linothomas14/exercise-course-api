package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/middleware"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/service"
)

type UserCourseController interface {
	FindAll(context *gin.Context)
	Create(context *gin.Context)
	GetUserCourseByID(context *gin.Context)
	Delete(context *gin.Context)
}

type userCourseController struct {
	userCourseService service.UserCourseService
}

func NewUserCourseController(userCourseService service.UserCourseService) UserCourseController {
	return &userCourseController{
		userCourseService: userCourseService,
	}
}

func (c *userCourseController) Create(ctx *gin.Context) {
	var reqParam param.UserCourseCreate

	err := ctx.ShouldBind(&reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = helper.ValidateStruct(reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if middleware.GetRoleFromClaims(ctx) != "admin" { // If "user" want to enroll, system will cross-check the id from token and req body must be same.

		idUserFromTokenTemp := middleware.GetUserIdFromClaims(ctx)
		idUserFromToken := uint32(idUserFromTokenTemp)

		if reqParam.UserID != idUserFromToken {

			res := helper.BuildResponse("You cant enroll other user profile", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

			return
		}

	}

	resp, err := c.userCourseService.CreateUserCourse(reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse("Enrolled", resp)
	ctx.JSON(http.StatusCreated, res)

}

func (c *userCourseController) GetUserCourseByID(ctx *gin.Context) {
	var userCourse model.UserCourse

	userCourseID := ctx.Param("id")

	ID, err := strconv.Atoi(userCourseID)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if ID == 0 {
		res := helper.BuildResponse("there is error occur", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userCourse, err = c.userCourseService.GetUserCourseByID(ID)

	parseUserCourse := parseUserCourseRes(userCourse)
	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse("OK", parseUserCourse)
	ctx.JSON(http.StatusOK, res)

}

func (c *userCourseController) FindAll(ctx *gin.Context) {
	userCourses, err := c.userCourseService.FindAll()

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse("OK", userCourses)
	ctx.JSON(http.StatusOK, res)
}

func (c *userCourseController) Delete(ctx *gin.Context) {
	var reqParam param.UserCourseCreate

	err := ctx.ShouldBind(&reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = helper.ValidateStruct(reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if middleware.GetRoleFromClaims(ctx) != "admin" { // If "user" want to enroll, system will cross-check the id from token and req body must be same.

		idUserFromTokenTemp := middleware.GetUserIdFromClaims(ctx)
		idUserFromToken := uint32(idUserFromTokenTemp)

		if reqParam.UserID != idUserFromToken {

			res := helper.BuildResponse("You cant enroll other user profile", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

			return
		}

	}

	err = c.userCourseService.Delete(reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(fmt.Sprintf("Success unenrolled user_id %d from course_id %d", reqParam.UserID, reqParam.CourseID), helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)

}

func parseUserCourseRes(userCourse model.UserCourse) response.UserCourseRes {

	res := response.UserCourseRes{ID: userCourse.ID,
		UserID: userCourse.UserID,
		User: response.UserResponse{
			ID:    userCourse.UserID,
			Email: userCourse.User.Email,
			Name:  userCourse.User.Name,
		},
		CourseId: userCourse.CourseID,
		Course: model.Course{
			ID:               userCourse.Course.ID,
			Title:            userCourse.Course.Title,
			CourseCategoryId: userCourse.Course.CourseCategoryId,
			CourseCategory: model.CourseCategory{
				ID:   userCourse.Course.CourseCategoryId,
				Name: userCourse.Course.CourseCategory.Name,
			},
		},
	}

	return res
}

func parseUserCourse(userCourseParam param.UserCourseCreate) model.UserCourse {
	var userCourse model.UserCourse

	userCourse.UserID = userCourseParam.UserID
	userCourse.CourseID = userCourseParam.CourseID

	return userCourse

}
