package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindById(id int) (User, error)
	FindAll() ([]User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(user User) (User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(id int) (User, error) {
	var user User

	err := r.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var user []User

	err := r.DB.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteUser(user User) (User, error) {
	err := r.DB.Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
