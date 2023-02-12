package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/service"
)

type UserController interface {
	GetProfile(context *gin.Context)
	Update(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetProfile(ctx *gin.Context) {
	var user response.UserResponse

	userID := helper.GetUserIdFromClaims(ctx)
	userRole := helper.GetRoleFromClaims(ctx)
	if userID == 0 {
		response := helper.BuildResponse("there is error occur", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.GetProfile(userID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", gin.H{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"role":  userRole,
	})

	ctx.JSON(http.StatusOK, response)

}

func (c *userController) Update(ctx *gin.Context) {

	var userParam param.UserUpdate
	var user model.User

	userID := helper.GetUserIdFromClaims(ctx)

	err := ctx.ShouldBind(&userParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(userParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user = parseUserUpdate(userParam)
	user.ID = uint32(userID)
	res, err := c.userService.Update(user)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("Updated", res)
	ctx.JSON(http.StatusOK, response)
}

func parseUserUpdate(userParam param.UserUpdate) model.User {
	var user model.User

	user.Name = userParam.Name
	user.Email = userParam.Email
	user.Password = userParam.Password

	return user

}
