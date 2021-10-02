package service

import (
	"context"
	"crudTaskEcho/internal/app/mocks"
	"crudTaskEcho/internal/app/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCRUDServiceImpl_CreateUser(t *testing.T) {
	mockRepo := new(mocks.CRUDRepository)

	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	service := NewCRUDService(mockRepo)

	err := service.CreateUser(context.Background(), &model.User{})
	assert.Nil(t, err)
}

func TestCRUDServiceImpl_GetUserById(t *testing.T) {
	mockRepo := new(mocks.CRUDRepository)

	mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(&model.User{}, nil)

	service := NewCRUDService(mockRepo)
	res, err := service.GetUserById(context.Background(), 2)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestCRUDServiceImpl_UpdateUser(t *testing.T) {
	mockRepo := new(mocks.CRUDRepository)

	mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return( nil)

	service := NewCRUDService(mockRepo)
	err := service.UpdateUser(context.Background(), &model.UpdateUserReq{})
	assert.Nil(t, err)
}