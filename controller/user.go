package controller

import (
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

type UserController interface {
	FindAll(context *gin.Context)
	GetProfile(context *gin.Context)
	GetUserByID(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
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

	userID := middleware.GetUserIdFromClaims(ctx)
	userRole := middleware.GetRoleFromClaims(ctx)
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
func (c *userController) GetUserByID(ctx *gin.Context) {
	var user response.UserResponse

	userID := ctx.Param("id")

	ID, err := strconv.Atoi(userID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if ID == 0 {
		response := helper.BuildResponse("there is error occur", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err = c.userService.GetProfile(ID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", user)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) Update(ctx *gin.Context) {

	var userParam param.UserUpdate
	var user model.User

	userID := middleware.GetUserIdFromClaims(ctx)

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

func (c *userController) FindAll(ctx *gin.Context) {
	users, err := c.userService.FindAll()

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", users)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	userId := uint32(u64)

	err = c.userService.Delete(userId)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("User id "+id+" was deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}

func parseUserUpdate(userParam param.UserUpdate) model.User {
	var user model.User

	user.Name = userParam.Name
	user.Email = userParam.Email
	user.Password = userParam.Password

	return user

}
