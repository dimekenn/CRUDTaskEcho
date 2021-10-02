package service

import (
	"context"
	"crudTaskEcho/internal/app/model"
)

type CRUDService interface {
	CreateUser(ctx context.Context, req *model.User) (*model.CreateUserRes, error)
	UpdateUser(ctx context.Context, req *model.UpdateUserReq) error
	GetUserById(ctx context.Context, id int) (*model.User, error)
}
