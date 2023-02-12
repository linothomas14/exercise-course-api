package controller

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/linothomas14/exercise-course-api/helper"
// 	"github.com/linothomas14/exercise-course-api/helper/response"
// 	"github.com/linothomas14/exercise-course-api/service"
// )

// type AdminController interface {
// 	GetProfile(context *gin.Context)
// 	// Update(context *gin.Context)
// }

// type adminController struct {
// 	adminService service.AdminService
// }

// func NewAdminController(adminService service.AdminService) AdminController {
// 	return &adminController{
// 		adminService: adminService,
// 	}
// }

// func (c *adminController) GetProfile(ctx *gin.Context) {
// 	var user response.UserResponse

// 	userID := helper.GetUserIdFromClaims(ctx)
// 	userRole := helper.GetRoleFromClaims(ctx)

// 	if userID == 0 || userRole == "" {
// 		response := helper.BuildResponse("there is error occur", helper.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	user, err := c.adminService.GetProfile(userID)

// 	if err != nil {
// 		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	response := helper.BuildResponse("OK", user)
// 	ctx.JSON(http.StatusOK, response)

// }
