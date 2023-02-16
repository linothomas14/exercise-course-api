package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/middleware"
	"github.com/linothomas14/exercise-course-api/service"
)

type AdminController interface {
	Register(context *gin.Context)
	GetProfile(context *gin.Context)
	GetAdminByID(context *gin.Context)
	FindAll(context *gin.Context)
}

type adminController struct {
	adminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &adminController{
		adminService: adminService,
	}
}

func (c *adminController) FindAll(ctx *gin.Context) {
	admins, err := c.adminService.FindAll()

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", admins)
	ctx.JSON(http.StatusOK, response)
}

func (c *adminController) GetProfile(ctx *gin.Context) {
	var user response.AdminResponse

	userID := middleware.GetUserIdFromClaims(ctx)

	if userID == 0 {
		response := helper.BuildResponse("there is error occur", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.adminService.GetProfile(userID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", user)
	ctx.JSON(http.StatusOK, response)

}
func (c *adminController) GetAdminByID(ctx *gin.Context) {
	var admin response.AdminResponse

	adminID := ctx.Param("id")

	ID, err := strconv.Atoi(adminID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	admin, err = c.adminService.GetProfile(ID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", admin)
	ctx.JSON(http.StatusOK, response)

}

func (c *adminController) Register(ctx *gin.Context) {
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

	createdAdmin, err := c.adminService.CreateAdmin(registerParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	registerResponse.ID = createdAdmin.ID
	registerResponse.Email = createdAdmin.Email
	registerResponse.Name = createdAdmin.Name
	response := helper.BuildResponse("OK", registerResponse)
	ctx.JSON(http.StatusCreated, response)
}
