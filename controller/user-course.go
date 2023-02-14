package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
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

	if middleware.GetRoleFromClaims(ctx) != "user" {
		res := helper.BuildResponse("You cant enroll class because you're admin, not a user", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctx.ShouldBind(&reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	idUserFromTokenTemp := middleware.GetUserIdFromClaims(ctx)
	idUserFromToken := uint32(idUserFromTokenTemp)

	reqParam.UserID = idUserFromToken

	err = helper.ValidateStruct(reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp, err := c.userCourseService.CreateUserCourse(reqParam)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse("Created", resp)
	ctx.JSON(http.StatusCreated, res)

}

func (c *userCourseController) GetUserCourseByID(ctx *gin.Context) {
	var userCourse *model.UserCourse

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

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse("OK", userCourse)
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

	id := ctx.Param("id")

	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	userCourseId := uint32(u64)

	err = c.userCourseService.Delete(userCourseId)

	if err != nil {
		res := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse("UserCourse id "+id+" was deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)

}

func parseUserCourse(userCourseParam param.UserCourseCreate) model.UserCourse {
	var userCourse model.UserCourse

	userCourse.UserID = userCourseParam.UserID
	userCourse.CourseID = userCourseParam.CourseID

	return userCourse

}
