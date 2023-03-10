package service

import (
	"fmt"

	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/repository"
)

type UserCourseService interface {
	FindAll() ([]response.UserCourseRes, error)
	GetUserCourseByID(int) (model.UserCourse, error)
	CreateUserCourse(userCourse param.UserCourseCreate) (*model.UserCourse, error)
	Delete(userCourse param.UserCourseCreate) error
}

type userCourseService struct {
	userCourseRepository repository.UserCourseRepository
}

func NewUserCourseService(userCourseRep repository.UserCourseRepository) UserCourseService {
	return &userCourseService{
		userCourseRepository: userCourseRep,
	}
}

func (service *userCourseService) FindAll() ([]response.UserCourseRes, error) {

	userCourse, err := service.userCourseRepository.FindAll()

	userCourses := parseFindAllUserCourse(userCourse)
	if err != nil {
		return []response.UserCourseRes{}, err
	}

	return userCourses, err
}

func (service *userCourseService) CreateUserCourse(userCourse param.UserCourseCreate) (*model.UserCourse, error) {

	userCourseToCreate := parseUserCourseParamToModel(userCourse)

	if service.userCourseRepository.IsDuplicateUserCourse(userCourseToCreate) {
		return nil, fmt.Errorf("You already enrolled on this course")
	}

	res, err := service.userCourseRepository.InsertUserCourse(userCourseToCreate)
	if err != nil {
		return userCourseToCreate, err
	}
	return res, err
}

func (service *userCourseService) GetUserCourseByID(ID int) (model.UserCourse, error) {

	res, err := service.userCourseRepository.GetUserCourseByID(ID)

	if err != nil {
		return model.UserCourse{}, err
	}

	return res, nil
}

func (service *userCourseService) Delete(userCourse param.UserCourseCreate) error {

	userCourseToCreate := parseUserCourseParamToModel(userCourse)

	if !service.userCourseRepository.IsDuplicateUserCourse(userCourseToCreate) {
		return fmt.Errorf("You not enrolled")
	}

	err := service.userCourseRepository.Delete(userCourseToCreate)
	if err != nil {
		return err
	}
	return err
}

// func parseFindAllUserCourse(userCourses []model.UserCourse) []response.UserCourseRes {
// 	var parsedUserCourse []response.UserCourseRes
// 	for _, userCourse := range userCourses {
// 		newCourse := response.UserCourseRes{
// 			ID:       userCourse.ID,
// 			UserID:   userCourse.UserID,
// 			CourseId: userCourse.CourseID,
// 		}
// 		parsedUserCourse = append(parsedUserCourse, newCourse)
// 	}
// 	return parsedUserCourse
// }

func parseUserCourseParamToModel(userParam param.UserCourseCreate) *model.UserCourse {
	var user model.UserCourse

	user.UserID = userParam.UserID
	user.CourseID = userParam.CourseID

	return &user

}

func parseFindAllUserCourse(userCourses []model.UserCourse) []response.UserCourseRes {

	var parsedUserCourse []response.UserCourseRes

	for _, userCourse := range userCourses {
		newUserCourse := response.UserCourseRes{
			ID:     userCourse.ID,
			UserID: userCourse.UserID,
			User: response.UserResponse{
				ID:    userCourse.UserID,
				Email: userCourse.User.Email,
				Name:  userCourse.User.Name,
			},
			CourseId: userCourse.CourseID,
			Course: model.Course{
				ID:               userCourse.Course.ID,
				Title:            userCourse.Course.Title,
				CourseCategoryId: userCourse.Course.CourseCategoryId,
				CourseCategory: model.CourseCategory{
					ID:   userCourse.Course.CourseCategoryId,
					Name: userCourse.Course.CourseCategory.Name,
				},
			},
		}
		parsedUserCourse = append(parsedUserCourse, newUserCourse)
	}
	return parsedUserCourse
}
