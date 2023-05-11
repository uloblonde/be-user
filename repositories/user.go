package repositories

import (
	"livecode/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	GetIdUSer(Id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, Id int) (models.User, error)
	DeleteUser(user models.User, Id int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) GetIdUSer(Id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, Id).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err

}
func (r *repository) UpdateUser(user models.User, Id int) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err

}

func (r *repository) DeleteUser(user models.User, Id int) (models.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
