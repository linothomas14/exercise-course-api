package service

import (
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/repository"
)

type CourseService interface {
	CreateCourse(model.Course) (model.Course, error)
	FindAll() ([]response.GetCoursesRes, error)
	FindByID(uint32) (model.Course, error)
	FindByName(string) (model.Course, error)
	Update(model.Course) (model.Course, error)
	Delete(uint32) error
}

type courseService struct {
	courseRepository repository.CourseRepository
}

func NewCourseService(courseRep repository.CourseRepository) CourseService {
	return &courseService{
		courseRepository: courseRep,
	}
}

func (service *courseService) FindAll() ([]response.GetCoursesRes, error) {

	course, err := service.courseRepository.FindAll()

	courses := parseFindAllCourses(course)
	if err != nil {
		return []response.GetCoursesRes{}, err
	}

	return courses, err
}

func (service *courseService) FindByID(id uint32) (model.Course, error) {

	course, err := service.courseRepository.FindByID(id)

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (service *courseService) FindByName(name string) (model.Course, error) {

	course, err := service.courseRepository.FindByName(name)

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (service *courseService) CreateCourse(course model.Course) (model.Course, error) {

	course, err := service.courseRepository.Insert(course)

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (service *courseService) Update(course model.Course) (model.Course, error) {

	course, err := service.courseRepository.Update(course)

	if err != nil {
		return model.Course{}, err
	}

	return course, err
}

func (service *courseService) Delete(id uint32) error {

	err := service.courseRepository.Delete(id)

	if err != nil {
		return err
	}

	return err
}

func parseFindAllCourses(courses []model.Course) []response.GetCoursesRes {
	var parsedCourses []response.GetCoursesRes
	for _, course := range courses {
		newCourse := response.GetCoursesRes{
			ID:               course.ID,
			Title:            course.Title,
			CourseCategoryId: course.CourseCategoryId,
		}
		parsedCourses = append(parsedCourses, newCourse)
	}
	return parsedCourses
}
