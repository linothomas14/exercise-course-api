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

type AdminController interface {
	Register(context *gin.Context)
	GetProfile(context *gin.Context)
	GetAdminByID(context *gin.Context)
	FindAll(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
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

func (c *adminController) Update(ctx *gin.Context) {

	var adminParam param.AdminUpdate
	var admin model.Admin

	id := ctx.Param("id")

	adminID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = ctx.ShouldBind(&adminParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(adminParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	adminParam.ID = uint32(adminID)

	admin = parseAdminUpdate(adminParam)

	res, err := c.adminService.Update(admin)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("Updated", res)
	ctx.JSON(http.StatusOK, response)
}
func (c *adminController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	adminId := uint32(u64)

	err = c.adminService.Delete(adminId)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("Course id "+id+" was deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}

func parseAdminUpdate(adminParam param.AdminUpdate) model.Admin {

	var admin model.Admin

	admin.ID = adminParam.ID
	admin.Name = adminParam.Name
	admin.Email = adminParam.Email
	admin.Password = adminParam.Password

	return admin

}
