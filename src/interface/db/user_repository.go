package db

import (
	"errors"

	"github.com/dj-hirrot/gorilla/src/domain/models"
)

type UserRepository struct {
	SqlHandler
}

func (repository *UserRepository) FindById(id int) (user models.User, err error) {
	if err = repository.First(&user, id).Error; err != nil {
		return
	}
	return
}

func (repository *UserRepository) FindAll() (users models.Users, err error) {
	if err = repository.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repository *UserRepository) Store(u models.User) (user models.User, err error) {
	if u.Name == "" {
		err = errors.New("*Name is required*")
		return
	}
	if u.Age < 0 {
		err = errors.New("*Age must be greater than 0*")
		return
	}
	if err = repository.Create(&u).Error; err != nil {
		return
	}
	return u, nil
}

func (repository *UserRepository) Update(u models.User) (user models.User, err error) {
	if u.Name == "" {
		err = errors.New("*Name is required*")
		return
	}
	if u.Age < 0 {
		err = errors.New("*Age must be greater than 0*")
		return
	}
	if err = repository.Save(&u).Error; err != nil {
		return
	}
	return u, nil
}

func (repository *UserRepository) DeleteById(u models.User) (err error) {
	if err = repository.Delete(&u).Error; err != nil {
		return
	}
	return
}
