package usecase

import (
	"github.com/dj-hirrot/gorilla/src/domain/entities"
	"github.com/dj-hirrot/gorilla/src/domain/models"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Index() (users models.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Show(id int) (user models.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Create(u models.User) (user models.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(id int, u entities.UserParams) (user models.User, err error) {
	user, err = interactor.UserRepository.Update(id, u)
	return
}

func (interactor *UserInteractor) Delete(u models.User) (err error) {
	err = interactor.UserRepository.DeleteById(u)
	return
}
