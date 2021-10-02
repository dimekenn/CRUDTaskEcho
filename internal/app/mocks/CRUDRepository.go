// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	model "crudTaskEcho/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// CRUDRepository is an autogenerated mock type for the CRUDRepository type
type CRUDRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, req
func (_m *CRUDRepository) CreateUser(ctx context.Context, req *model.User) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *CRUDRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, req
func (_m *CRUDRepository) UpdateUser(ctx context.Context, req *model.UpdateUserReq) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdateUserReq) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
