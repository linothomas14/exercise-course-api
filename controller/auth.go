package controller

import (
	"net/http"

	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(authService service.AuthService, userService service.UserService) AuthController {
	return &authController{
		authService: authService,
		userService: userService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginparam param.Login
	var role string

	err := ctx.ShouldBind(&loginparam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(loginparam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	path := ctx.Request.URL.Path

	if path == "/login-admin" {
		role = "admin" //this for admin
	} else {
		role = "user" //this for user
	}
	token, err := c.authService.Login(loginparam.Email, loginparam.Password, role)

	if token == "" {
		response := helper.BuildResponse("invalid credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", gin.H{
		"token": token,
	})
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *authController) Register(ctx *gin.Context) {
	var registerParam param.Register
	var registerResponse response.RegisterResponse

	err := ctx.ShouldBind(&registerParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(registerParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)

		return
	}

	if c.authService.IsDuplicateEmail(registerParam.Email) {
		response := helper.BuildResponse("Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	createdUser, err := c.userService.CreateUser(registerParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	registerResponse.ID = createdUser.ID
	registerResponse.Email = createdUser.Email
	registerResponse.Name = createdUser.Name
	response := helper.BuildResponse("OK", registerResponse)
	ctx.JSON(http.StatusCreated, response)
}
