package main

import (
	"time"

	"github.com/linothomas14/exercise-course-api/config"
	"github.com/linothomas14/exercise-course-api/controller"
	"github.com/linothomas14/exercise-course-api/middleware"
	"github.com/linothomas14/exercise-course-api/repository"
	"github.com/linothomas14/exercise-course-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepository           repository.UserRepository           = repository.NewUserRepository(db)
	courseCategoryRepository repository.CourseCategoryRepository = repository.NewCourseCategoryRepository(db)
	courseRepository         repository.CourseRepository         = repository.NewCourseRepository(db)
	adminRepository          repository.AdminRepository          = repository.NewAdminRepository(db)
	userCourseRepository     repository.UserCourseRepository     = repository.NewUserCourseRepository(db)

	userService           service.UserService           = service.NewUserService(userRepository)
	authService           service.AuthService           = service.NewAuthService(userRepository, adminRepository)
	courseCategoryService service.CourseCategoryService = service.NewCourseCategoryService(courseCategoryRepository)
	courseService         service.CourseService         = service.NewCourseService(courseRepository)
	adminService          service.AdminService          = service.NewAdminService(adminRepository)
	userCourseService     service.UserCourseService     = service.NewUserCourseService(userCourseRepository)

	authController           controller.AuthController           = controller.NewAuthController(authService, userService)
	userController           controller.UserController           = controller.NewUserController(userService)
	courseCategoryController controller.CourseCategoryController = controller.NewCourseCategoryController(courseCategoryService)
	courseController         controller.CourseController         = controller.NewCourseController(courseService)
	adminController          controller.AdminController          = controller.NewAdminController(adminService)
	userCourseController     controller.UserCourseController     = controller.NewUserCourseController(userCourseService)
)

func PingHandler(c *gin.Context) {
	t := time.Now()
	c.JSON(200, gin.H{
		"msg":  "pong",
		"time": t,
	})
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.POST("/login", authController.Login) //login for user
	r.POST("/register", authController.Register)
	r.POST("/login-admin", authController.Login)

	adminRoutes := r.Group("admins", middleware.AuthorizeJWT(), middleware.AuthorizeRole([]string{"admin"}))
	{

		adminRoutes.POST("/", adminController.Register) //register new admin
		adminRoutes.GET("/", adminController.FindAll)
		adminRoutes.GET("/:id", adminController.GetAdminByID)
		adminRoutes.PUT("/:id", PingHandler)
		adminRoutes.DELETE("/:id", PingHandler)
	}
	userRoutes := r.Group("users", middleware.AuthorizeJWT())
	{
		userRoutes.GET("/", middleware.AuthorizeRole([]string{"admin"}), userController.FindAll)
		userRoutes.GET("/:id", middleware.AuthorizeRole([]string{"admin"}), userController.GetUserByID)
		userRoutes.GET("/profile", middleware.AuthorizeRole([]string{"user"}), userController.GetProfile)
		userRoutes.PUT("/", middleware.AuthorizeRole([]string{"user"}), userController.Update)
		userRoutes.DELETE("/:id", middleware.AuthorizeRole([]string{"admin"}), userController.Delete)
	}

	courseRoutes := r.Group("courses", middleware.AuthorizeJWT())
	{
		courseRoutes.GET("/", courseController.FindAll)
		courseRoutes.GET("/:id", courseController.FindByID)
		courseRoutes.POST("/", middleware.AuthorizeRole([]string{"admin"}), courseController.Create)
		courseRoutes.PUT("/:id", middleware.AuthorizeRole([]string{"admin"}), PingHandler)
		courseRoutes.DELETE("/:id", middleware.AuthorizeRole([]string{"admin"}), PingHandler)
	}

	attendanceRoutes := r.Group("user-courses", middleware.AuthorizeJWT())
	{
		attendanceRoutes.GET("/", middleware.AuthorizeRole([]string{"admin"}), userCourseController.FindAll)
		attendanceRoutes.POST("/", userCourseController.Create)   //Enroll Course
		attendanceRoutes.DELETE("/", userCourseController.Delete) //delete their course, cant delete other user course
	}

	CourseCategoryRoutes := r.Group("course-categories")
	{
		CourseCategoryRoutes.GET("/", courseCategoryController.FindAll)
		CourseCategoryRoutes.GET("/:id", courseCategoryController.FindByID)
		CourseCategoryRoutes.POST("/", middleware.AuthorizeRole([]string{"admin"}), courseCategoryController.Create)
		CourseCategoryRoutes.PUT("/:id", middleware.AuthorizeRole([]string{"admin"}), courseCategoryController.Update)
		CourseCategoryRoutes.DELETE("/:id", middleware.AuthorizeRole([]string{"admin"}), courseCategoryController.Delete)
	}

	r.GET("ping", PingHandler)
	r.Run()
}
