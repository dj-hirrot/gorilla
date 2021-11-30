package mock_usecase

import (
	"reflect"
	"testing"

	"github.com/dj-hirrot/gorilla/src/domain/models"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
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

	id := uuid.NewV4()

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

	userAttributes := models.UserAttributes{
		Name: "阿南惟幾",
	}
	user := models.User{
		UserAttributes: userAttributes,
	}
	var expected models.User
	var err error

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Store(user).Return(expected, err)

	result, err := mockUserRepository.Store(user)

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

	id := uuid.NewV4()
	userAttributes := models.UserAttributes{
		Name: "東郷平八郎",
	}
	user := models.User{
		UserAttributes: userAttributes,
	}
	var expected models.User
	var err error

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Update(id, user).Return(expected, err)

	result, err := mockUserRepository.Update(id, user)

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

	id := uuid.NewV4()

	mockUserRepository := NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().DeleteById(id).Return(err)

	err = mockUserRepository.DeleteById(id)

	if err != nil {
		t.Errorf("Expected nil, got %v\n", err)
	}
}
