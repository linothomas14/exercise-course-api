package repository

import (
	"github.com/linothomas14/exercise-course-api/model"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Insert(model.Course) (model.Course, error)
	FindByID(uint32) (model.Course, error)
	FindByName(string) (model.Course, error)
	FindAll() ([]model.Course, error)
	Update(model.Course) (model.Course, error)
	Delete(uint32) error
}

type courseConnection struct {
	connection *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseConnection{
		connection: db,
	}
}

func (db *courseConnection) Insert(course model.Course) (model.Course, error) {

	err := db.connection.Preload("CourseCategory").Save(&course).Find(&course).Error

	if err != nil {
		return model.Course{}, err
	}
	return course, err
}

func (db *courseConnection) FindAll() ([]model.Course, error) {

	var course []model.Course

	err := db.connection.Preload("CourseCategory").Find(&course).Error

	if err != nil {
		return []model.Course{}, err
	}

	return course, err
}

func (db *courseConnection) FindByID(id uint32) (model.Course, error) {

	var course model.Course

	course.ID = id

	err := db.connection.Preload("CourseCategory").First(&course).Error

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (db *courseConnection) FindByName(name string) (model.Course, error) {

	var course model.Course

	err := db.connection.Preload("CourseCategory").Where("name like ?", "%"+name+"%").Find(&course).Error

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (db *courseConnection) Update(course model.Course) (model.Course, error) {

	err := db.connection.Model(&course).Updates(&course).Find(&course).Error

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (db *courseConnection) Delete(id uint32) error {

	var course model.Course

	course.ID = id
	err := db.connection.Delete(&course).Error

	return err
}
