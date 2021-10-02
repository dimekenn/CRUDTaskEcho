package repository

import (
	"context"
	"crudTaskEcho/internal/app/mocks"
	"crudTaskEcho/internal/app/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCRUDRepositoryImpl_CreateUser(t *testing.T) {
	repoMock := new(mocks.CRUDRepository)
	repoMock.On("CreateUser", mock.Anything, mock.Anything).Return(nil)
	err := repoMock.CreateUser(context.Background(), &model.User{})
	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}

func TestCRUDRepositoryImpl_GetUserById(t *testing.T) {
	repoMock := new(mocks.CRUDRepository)
	repoMock.On("GetUserById", mock.Anything, mock.Anything).Return(&model.User{}, nil)
	res, err := repoMock.GetUserById(context.Background(), 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	repoMock.AssertExpectations(t)
}

func TestCRUDRepositoryImpl_UpdateUser(t *testing.T) {
	repoMock := new(mocks.CRUDRepository)
	repoMock.On("UpdateUser", mock.Anything, mock.Anything).Return(nil)
	err := repoMock.UpdateUser(context.Background(), &model.UpdateUserReq{})
	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}
