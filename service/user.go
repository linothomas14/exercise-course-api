package service

import (
	"log"

	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	FindAll() ([]response.UserResponse, error)
	CreateUser(user param.Register) (model.User, error)
	FindByEmail(email string) model.User
	Update(user model.User) (response.UserResponse, error)
	GetProfile(userId int) (response.UserResponse, error)
	Delete(userId uint32) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

func (service *userService) FindAll() ([]response.UserResponse, error) {

	user, err := service.userRepository.FindAll()

	users := parseFindAllUser(user)
	if err != nil {
		return []response.UserResponse{}, err
	}

	return users, err
}

func (service *userService) CreateUser(user param.Register) (model.User, error) {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return userToCreate, err
	}
	res, err := service.userRepository.InsertUser(userToCreate)
	if err != nil {
		return userToCreate, err
	}
	return res, err
}

func (service *userService) Update(userParam model.User) (response.UserResponse, error) {
	var userRes response.UserResponse

	if userParam.Password != "" {
		userParam.Password = helper.HashAndSalt([]byte(userParam.Password))
	}

	user, err := service.userRepository.UpdateUser(userParam)

	if err != nil {
		return response.UserResponse{}, err
	}

	userRes.ID = user.ID
	userRes.Name = user.Name
	userRes.Email = user.Email

	return userRes, nil
}

func (service *userService) GetProfile(userId int) (response.UserResponse, error) {

	var userRes response.UserResponse

	user, err := service.userRepository.GetUser(userId)
	if err != nil {
		return response.UserResponse{}, err
	}

	userRes.ID = user.ID
	userRes.Name = user.Name
	userRes.Email = user.Email

	return userRes, err
}

func (service *userService) FindByEmail(email string) model.User {
	return service.userRepository.FindByEmail(email)
}

func (service *userService) Delete(id uint32) error {

	err := service.userRepository.Delete(id)

	return err
}

func parseFindAllUser(users []model.User) []response.UserResponse {
	var parsedUser []response.UserResponse
	for _, user := range users {
		newCourse := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		parsedUser = append(parsedUser, newCourse)
	}
	return parsedUser
}
