package main

import (
	"time"

	// "github.com/linothomas14/exercise-course-api/controller"
	// "github.com/linothomas14/exercise-course-api/middleware"
	"github.com/linothomas14/exercise-course-api/controller"
	"github.com/linothomas14/exercise-course-api/repository"

	// "github.com/linothomas14/exercise-course-api/service"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/exercise-course-api/config"
	"github.com/linothomas14/exercise-course-api/middleware"
	"github.com/linothomas14/exercise-course-api/service"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepository           repository.UserRepository           = repository.NewUserRepository(db)
	courseCategoryRepository repository.CourseCategoryRepository = repository.NewCourseCategoryRepository(db)
	courseRepository         repository.CourseRepository         = repository.NewCourseRepository(db)
	adminRepository          repository.AdminRepository          = repository.NewAdminRepository(db)
	// // 	transactionRepository  repository.TransactionRepository  = repository.NewTransactionRepository(db)
	// // jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)
	authService service.AuthService = service.NewAuthService(userRepository, adminRepository)

	courseCategoryService service.CourseCategoryService = service.NewCourseCategoryService(courseCategoryRepository)
	courseService         service.CourseService         = service.NewCourseService(courseRepository)
	// attendanceService service.AttendanceService = service.NewAttendanceService(attendanceRepository, userRepository, courseCategoryRepository)

	// // 	transactionService  service.TransactionService  = service.NewTransactionService(transactionRepository, productRepository)

	authController           controller.AuthController           = controller.NewAuthController(authService, userService)
	userController           controller.UserController           = controller.NewUserController(userService)
	courseCategoryController controller.CourseCategoryController = controller.NewCourseCategoryController(courseCategoryService)
	courseController         controller.CourseController         = controller.NewCourseController(courseService)
	// attendanceController controller.AttendanceController = controller.NewAttendanceController(attendanceService)
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

	authRoutes := r.Group("admins", middleware.AuthorizeJWT(authService))
	{

		authRoutes.POST("/", PingHandler) //register new admin
		authRoutes.GET("/", PingHandler)
		authRoutes.GET("/:id", PingHandler)
		authRoutes.PUT("/:id", PingHandler)
		authRoutes.DELETE("/:id", PingHandler)
	}

	userRoutes := r.Group("users", middleware.AuthorizeJWT(authService))
	{
		userRoutes.GET("/", userController.GetProfile)
		userRoutes.PUT("/", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	courseRoutes := r.Group("courses")
	{
		courseRoutes.GET("/", courseController.FindAll)
		courseRoutes.GET("/:id", courseController.FindByID)
		courseRoutes.POST("/", courseController.Create)
		courseRoutes.PUT("/:id", PingHandler)
		courseRoutes.DELETE("/:id", PingHandler)
	}
	attendanceRoutes := r.Group("user-course", middleware.AuthorizeJWT(authService))
	{
		attendanceRoutes.GET("/", PingHandler)
		attendanceRoutes.POST("/", PingHandler)
		attendanceRoutes.DELETE("/:idUserCourse", PingHandler)
	}

	CourseCategoryRoutes := r.Group("course-category")
	{
		CourseCategoryRoutes.GET("/", courseCategoryController.FindAll)
		CourseCategoryRoutes.GET("/:id", courseCategoryController.FindByID)
		CourseCategoryRoutes.POST("/", courseCategoryController.Create)
		CourseCategoryRoutes.PUT("/:id", courseCategoryController.Update)
		CourseCategoryRoutes.DELETE("/:id", courseCategoryController.Delete)
	}

	r.GET("ping", PingHandler)
	r.Run()
}
