package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	// .ENV VAR
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// DNS WITH PORT
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	// DNS WITHOUT PORT
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn+"&loc=Asia%2FJakarta"), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to DB")
	}

	db.Debug().AutoMigrate(&model.Course{}, &model.User{}, &model.Admin{}, &model.UserCourse{}, &model.CourseCategory{})
	Seeders(db)
	fmt.Println("Successfully connected!")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection to DB")
	}
	dbSQL.Close()
}

func Seeders(db *gorm.DB) {
	var user model.Admin = model.Admin{
		Email:    "admin@gmail.com",
		Name:     "Admin",
		Password: helper.HashAndSalt([]byte("admin123")),
	}

	err := db.First(&user, "email = ?", user.Email).Error
	if err != nil {
		db.Create(&user)
	}
}
