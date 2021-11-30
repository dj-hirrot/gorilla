package db

import (
	"errors"

	"github.com/dj-hirrot/gorilla/src/domain/models"
	uuid "github.com/satori/go.uuid"
)

type UserRepository struct {
	SqlHandler
}

func (repository *UserRepository) FindById(id uuid.UUID) (user models.User, err error) {
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
	if err = repository.Create(&u).Error; err != nil {
		return
	}
	return u, nil
}

func (repository *UserRepository) Update(id uuid.UUID, u models.User) (user models.User, err error) {
	if u.Name == "" {
		err = errors.New("*Name is required*")
		return
	}
	if err = repository.First(&user, id).Error; err != nil {
		return
	}
	user.Name = u.Name
	if err = repository.Save(&user).Error; err != nil {
		return
	}
	return
}

func (repository *UserRepository) DeleteById(id uuid.UUID) (err error) {
	u := models.User{}
	if err = repository.First(&u, id).Error; err != nil {
		return
	}
	if err = repository.Delete(&u).Error; err != nil {
		return
	}
	return
}
