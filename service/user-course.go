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
	GetUserCourseByID(int) (*model.UserCourse, error)
	CreateUserCourse(userCourse param.UserCourseCreate) (*model.UserCourse, error)
	Update(userCourse model.UserCourse) (response.UserCourseRes, error)
	Delete(userCourseId uint32) error
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

	userCourseToCreate := parseUserCourse(userCourse)

	if service.userCourseRepository.IsDuplicateUserCourse(userCourseToCreate) {
		return nil, fmt.Errorf("You already enrolled on this course")
	}

	res, err := service.userCourseRepository.InsertUserCourse(userCourseToCreate)
	if err != nil {
		return userCourseToCreate, err
	}
	return res, err
}

func (service *userCourseService) GetUserCourseByID(ID int) (*model.UserCourse, error) {

	res, err := service.GetUserCourseByID(ID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *userCourseService) Update(userCourseParam model.UserCourse) (response.UserCourseRes, error) {
	var userCourseRes response.UserCourseRes

	userCourse, err := service.userCourseRepository.UpdateUserCourse(userCourseParam)

	if err != nil {
		return response.UserCourseRes{}, err
	}

	userCourseRes.ID = userCourse.ID
	userCourseRes.UserID = userCourse.UserID
	userCourseRes.CourseId = userCourse.CourseID

	return userCourseRes, nil
}

func (service *userCourseService) Delete(id uint32) error {

	err := service.userCourseRepository.Delete(id)

	return err
}

func parseFindAllUserCourse(userCourses []model.UserCourse) []response.UserCourseRes {
	var parsedUserCourse []response.UserCourseRes
	for _, userCourse := range userCourses {
		newCourse := response.UserCourseRes{
			ID:       userCourse.ID,
			UserID:   userCourse.UserID,
			CourseId: userCourse.CourseID,
		}
		parsedUserCourse = append(parsedUserCourse, newCourse)
	}
	return parsedUserCourse
}

func parseUserCourse(userParam param.UserCourseCreate) *model.UserCourse {
	var user model.UserCourse

	user.UserID = userParam.UserID
	user.CourseID = userParam.CourseID

	return &user

}
