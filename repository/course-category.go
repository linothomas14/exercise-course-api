package repository

import (
	"github.com/linothomas14/exercise-course-api/model"
	"gorm.io/gorm"
)

type CourseCategoryRepository interface {
	Insert(CourseCategory model.CourseCategory) (model.CourseCategory, error)
	FindByID(uint32) (model.CourseCategory, error)
	FindAll() ([]model.CourseCategory, error)
	Update(CourseCategory model.CourseCategory) (model.CourseCategory, error)
	Delete(uint32) error
}

type courseCategoryConnection struct {
	connection *gorm.DB
}

func NewCourseCategoryRepository(db *gorm.DB) CourseCategoryRepository {
	return &courseCategoryConnection{
		connection: db,
	}
}

func (db *courseCategoryConnection) Insert(courseCategory model.CourseCategory) (model.CourseCategory, error) {

	err := db.connection.Create(&courseCategory).Error

	if err != nil {
		return model.CourseCategory{}, err
	}

	return courseCategory, err
}

func (db *courseCategoryConnection) FindByID(id uint32) (model.CourseCategory, error) {

	var ctg model.CourseCategory

	err := db.connection.First(&ctg, "id=?", id).Error

	if err != nil {
		return model.CourseCategory{}, err
	}

	return ctg, err
}

func (db *courseCategoryConnection) FindAll() ([]model.CourseCategory, error) {

	var ctg []model.CourseCategory

	err := db.connection.Find(&ctg).Error

	if err != nil {
		return []model.CourseCategory{}, err
	}

	return ctg, err
}

func (db *courseCategoryConnection) Update(ctg model.CourseCategory) (model.CourseCategory, error) {

	err := db.connection.Model(&ctg).Updates(&ctg).Find(&ctg).Error

	if err != nil {
		return model.CourseCategory{}, err
	}

	return ctg, err
}

func (db *courseCategoryConnection) Delete(id uint32) error {

	var ctg model.CourseCategory

	ctg.ID = id
	err := db.connection.Delete(&ctg).Error

	return err

}
