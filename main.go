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
	// attendanceRepository repository.AttendanceRepository = repository.NewAttendanceRepository(db)
	// // 	transactionRepository  repository.TransactionRepository  = repository.NewTransactionRepository(db)
	// // jwtService  service.JWTService  = service.NewJWTService()
	// userService       service.UserService       = service.NewUserService(userRepository)
	// authService       service.AuthService       = service.NewAuthService(userRepository)
	jwtService            service.AuthService           = service.NewAuthService(userRepository)
	courseCategoryService service.CourseCategoryService = service.NewCourseCategoryService(courseCategoryRepository)
	courseService         service.CourseService         = service.NewCourseService(courseRepository)
	// attendanceService service.AttendanceService = service.NewAttendanceService(attendanceRepository, userRepository, courseCategoryRepository)

	// // 	transactionService  service.TransactionService  = service.NewTransactionService(transactionRepository, productRepository)

	// authController       controller.AuthController       = controller.NewAuthController(authService)
	// userController       controller.UserController       = controller.NewUserController(userService)
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

	r.POST("/login", PingHandler)
	r.POST("/register", PingHandler)

	authRoutes := r.Group("admins", middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.POST("/", PingHandler)
		authRoutes.GET("/:idAdmin", PingHandler)
		authRoutes.PUT("/:idAdmin", PingHandler)
		authRoutes.DELETE("/:idAdmin", PingHandler)
	}

	userRoutes := r.Group("users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", PingHandler)
		userRoutes.GET("/:idUser", PingHandler)
		userRoutes.PUT("/:idUser", PingHandler)
		userRoutes.DELETE("/delete", PingHandler)
	}

	courseRoutes := r.Group("courses")
	{
		courseRoutes.GET("/", courseController.FindAll)
		courseRoutes.GET("/:id", courseController.FindByID)
		courseRoutes.POST("/", courseController.Create)
		courseRoutes.PUT("/:id", PingHandler)
		courseRoutes.DELETE("/:id", PingHandler)
	}
	attendanceRoutes := r.Group("user-course", middleware.AuthorizeJWT(jwtService))
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
