package service

import (
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/helper/param"
	"github.com/linothomas14/exercise-course-api/helper/response"
	"github.com/linothomas14/exercise-course-api/model"
	"github.com/linothomas14/exercise-course-api/repository"
)

type AdminService interface {
	CreateAdmin(admin param.Register) (*model.Admin, error)
	FindAll() ([]response.AdminResponse, error)
	FindByEmail(email string) model.Admin
	Update(admin model.Admin) (response.AdminResponse, error)
	GetProfile(adminId int) (response.AdminResponse, error)
	Delete(adminId uint32) error
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(adminRep repository.AdminRepository) AdminService {
	return &adminService{
		adminRepository: adminRep,
	}
}

func (service *adminService) FindAll() ([]response.AdminResponse, error) {

	res, err := service.adminRepository.FindAll()

	if err != nil {
		return nil, err
	}

	parseAdmin := parseAdmins(res)

	return parseAdmin, nil
}

func (service *adminService) CreateAdmin(admin param.Register) (*model.Admin, error) {

	adminToCreate := &model.Admin{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}

	res, err := service.adminRepository.InsertAdmin(adminToCreate)
	if err != nil {
		return adminToCreate, err
	}
	return res, err
}

func (service *adminService) Update(adminParam model.Admin) (response.AdminResponse, error) {
	var adminRes response.AdminResponse

	if adminParam.Password != "" {
		adminParam.Password = helper.HashAndSalt([]byte(adminParam.Password))
	}

	admin, err := service.adminRepository.UpdateAdmin(adminParam)

	if err != nil {
		return response.AdminResponse{}, err
	}

	adminRes.ID = admin.ID
	adminRes.Name = admin.Name
	adminRes.Email = admin.Email

	return adminRes, nil
}

func (service *adminService) GetProfile(adminId int) (response.AdminResponse, error) {

	var adminRes response.AdminResponse

	admin, err := service.adminRepository.GetAdmin(adminId)
	if err != nil {
		return response.AdminResponse{}, err
	}

	adminRes.ID = admin.ID
	adminRes.Name = admin.Name
	adminRes.Email = admin.Email

	return adminRes, err
}

func (service *adminService) FindByEmail(email string) model.Admin {
	return service.adminRepository.FindByEmail(email)
}

func (service *adminService) Delete(id uint32) error {

	err := service.adminRepository.Delete(id)

	return err
}

func parseAdmins(admins []model.Admin) []response.AdminResponse {
	var parsedAdmin []response.AdminResponse

	for _, admin := range admins {
		newAdmin := response.AdminResponse{
			ID:    admin.ID,
			Name:  admin.Name,
			Email: admin.Email,
		}
		parsedAdmin = append(parsedAdmin, newAdmin)
	}
	return parsedAdmin
}
