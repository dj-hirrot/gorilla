package mock_usecase

import (
	"reflect"
	"testing"

	"github.com/dj-hirrot/gorilla/src/domain/entities"
	"github.com/dj-hirrot/gorilla/src/domain/models"
	"github.com/golang/mock/gomock"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected models.Users
	var err error

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindAll().Return(expected, err)

	result, err := mockUserRepository.FindAll()

	if err != nil {
		t.Errorf("Expected nil, got %v\n", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected nil, got %v\n", err)
	}
}

func TestShow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected models.User
	var err error

	id := 1

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindById(id).Return(expected, err)

	result, err := mockUserRepository.FindById(id)

	if err != nil {
		t.Errorf("Expected nil, got %v\n", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected nil, got %v\n", err)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := models.User{
		Name: "阿南惟幾",
		Age:  58,
	}
	var err error

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Store(expected).Return(expected, err)

	result, err := mockUserRepository.Store(expected)

	if err != nil {
		t.Errorf("Expected nil, got %v\n", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected nil, got %v\n", err)
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := 1
	body := entities.UserParams{
		Name: "東郷平八郎",
		Age:  86,
	}
	expected := models.User{
		Id:   id,
		Name: "東郷平八郎",
		Age:  86,
	}
	var err error

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Update(id, body).Return(expected, err)

	result, err := mockUserRepository.Update(id, body)

	if err != nil {
		t.Errorf("Expected nil, got %v\n", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected nil, got %v\n", err)
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error

	id := 1

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().DeleteById(id).Return(err)

	err = mockUserRepository.DeleteById(id)

	if err != nil {
		t.Errorf("Expected nil, got %v\n", err)
	}
}
