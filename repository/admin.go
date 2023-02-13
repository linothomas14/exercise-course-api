package repository

import (
	"github.com/linothomas14/exercise-course-api/helper"
	"github.com/linothomas14/exercise-course-api/model"
	"gorm.io/gorm"
)

//AdminRepository is contract what adminRepository can do to db
type AdminRepository interface {
	InsertAdmin(admin *model.Admin) (*model.Admin, error)
	UpdateAdmin(admin model.Admin) (model.Admin, error)
	VerifyCredential(email string, password string) interface{}
	FindByEmail(email string) model.Admin
	GetAdmin(adminID int) (model.Admin, error)
	Delete(adminID uint32) error
}

type adminConnection struct {
	connection *gorm.DB
}

//NewAdminRepository is creates a new instance of AdminRepository
func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminConnection{
		connection: db,
	}
}

func (db *adminConnection) InsertAdmin(admin *model.Admin) (*model.Admin, error) {

	admin.Password = helper.HashAndSalt([]byte(admin.Password))
	err := db.connection.Save(&admin).Error
	return admin, err
}

func (db *adminConnection) UpdateAdmin(admin model.Admin) (model.Admin, error) {

	err := db.connection.Model(&admin).Updates(&admin).Find(&admin).Error
	return admin, err
}

func (db *adminConnection) VerifyCredential(email string, password string) interface{} {
	var admin model.Admin
	res := db.connection.Where("email = ?", email).Take(&admin)
	if res.Error == nil {
		return admin
	}
	return nil
}

func (db *adminConnection) FindByEmail(email string) model.Admin {
	var admin model.Admin
	db.connection.Where("email = ?", email).Take(&admin)
	return admin
}

func (db *adminConnection) GetAdmin(adminId int) (model.Admin, error) {
	var admin model.Admin
	err := db.connection.Find(&admin, adminId).Error
	return admin, err
}

func (db *adminConnection) Delete(adminID uint32) error {
	var admin model.Admin
	admin.ID = adminID
	err := db.connection.First(&admin).Error

	if err != nil {
		return err
	}

	err = db.connection.Delete(&admin).Error

	return err
}
