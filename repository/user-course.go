package repository

import (
	"github.com/linothomas14/exercise-course-api/model"
	"gorm.io/gorm"
)

//UserCourseRepository is contract what userCourseRepository can do to db
type UserCourseRepository interface {
	FindAll() ([]model.UserCourse, error)
	IsDuplicateUserCourse(userCourse *model.UserCourse) bool
	GetUserCourseByID(userCourseID int) (model.UserCourse, error)
	InsertUserCourse(userCourse *model.UserCourse) (*model.UserCourse, error)

	Delete(userCourse *model.UserCourse) error
}

type userCourseConnection struct {
	connection *gorm.DB
}

//NewUserCourseRepository is creates a new instance of UserCourseRepository
func NewUserCourseRepository(db *gorm.DB) UserCourseRepository {
	return &userCourseConnection{
		connection: db,
	}
}
func (db *userCourseConnection) FindAll() ([]model.UserCourse, error) {

	var userCourse []model.UserCourse

	err := db.connection.Preload("Course.CourseCategory").Preload("User").Find(&userCourse).Error

	if err != nil {
		return []model.UserCourse{}, err
	}

	return userCourse, err
}

func (db *userCourseConnection) InsertUserCourse(userCourse *model.UserCourse) (*model.UserCourse, error) {

	err := db.connection.Preload("Course.CourseCategory").Preload("User").Save(&userCourse).Find(&userCourse).Error

	return userCourse, err
}

func (db *userCourseConnection) IsDuplicateUserCourse(userCourse *model.UserCourse) bool {

	err := db.connection.Table("user_course").Where("user_id = ? AND course_id = ? ", userCourse.UserID, userCourse.CourseID).First(&userCourse).Error

	if err != nil {
		if err.Error() == "record not found" {
			return false
		}
	}

	return true
}
func (db *userCourseConnection) GetUserCourseByID(userCourseId int) (model.UserCourse, error) {
	var userCourse model.UserCourse
	err := db.connection.Preload("Course.CourseCategory").Preload("User").First(&userCourse, userCourseId).Error

	return userCourse, err
}

func (db *userCourseConnection) Delete(userCourse *model.UserCourse) error {

	err := db.connection.Where("user_id = ? AND course_id = ? ", userCourse.UserID, userCourse.CourseID).Delete(&userCourse).Error

	return err
}
