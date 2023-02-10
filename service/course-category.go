package service

import (
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/repository"
)

type CourseCategoryService interface {
	CreateCourseCategory(model.CourseCategory) (model.CourseCategory, error)
	FindAll() ([]model.CourseCategory, error)
	FindByID(uint32) (model.CourseCategory, error)
	Update(model.CourseCategory) (model.CourseCategory, error)
	Delete(uint32) error
}

type courseCategoryService struct {
	courseCategoryRepository repository.CourseCategoryRepository
}

func NewCourseCategoryService(courseCategoryRep repository.CourseCategoryRepository) CourseCategoryService {
	return &courseCategoryService{
		courseCategoryRepository: courseCategoryRep,
	}
}

func (service *courseCategoryService) CreateCourseCategory(courseCategory model.CourseCategory) (model.CourseCategory, error) {

	courseCategory, err := service.courseCategoryRepository.Insert(courseCategory)

	if err != nil {
		return model.CourseCategory{}, err
	}

	return courseCategory, err
}

func (service *courseCategoryService) FindAll() ([]model.CourseCategory, error) {

	courseCategory, err := service.courseCategoryRepository.FindAll()

	if err != nil {
		return []model.CourseCategory{}, err
	}

	return courseCategory, err
}

func (service *courseCategoryService) FindByID(id uint32) (model.CourseCategory, error) {

	courseCategory, err := service.courseCategoryRepository.FindByID(id)

	if err != nil {
		return model.CourseCategory{}, err
	}

	return courseCategory, err
}

func (service *courseCategoryService) Update(ctg model.CourseCategory) (model.CourseCategory, error) {

	ctg, err := service.courseCategoryRepository.Update(ctg)

	if err != nil {
		return model.CourseCategory{}, err
	}

	return ctg, err
}

func (service *courseCategoryService) Delete(id uint32) error {

	err := service.courseCategoryRepository.Delete(id)

	return err
}
